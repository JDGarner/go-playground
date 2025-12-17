package ratelimiter

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// TODO:
// - replace impl with use redis

type RateLimiter struct {
	mu         sync.Mutex
	limit      int
	timeWindow time.Duration
	client     *redis.Client
}

type OptionFunc func(o *RateLimiter)

func WithLimit(limit int) OptionFunc {
	return func(o *RateLimiter) {
		o.limit = limit
	}
}

func New(opts ...OptionFunc) *RateLimiter {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		DialTimeout:  1 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		PoolSize:     10,
	})

	rl := &RateLimiter{
		limit:      10,
		timeWindow: 1 * time.Minute,
		client:     client,
	}

	for _, opt := range opts {
		opt(rl)
	}

	return rl
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIPAddress(r)
		key := rl.GetRequestKey(r.URL.Path, ip)

		allowed, remaining, resetAt, err := rl.CheckRateLimit(r.Context(), key)
		if err != nil {
			fmt.Println(">>> err: ", err)
		}

		fmt.Printf(">>> allowed: %v, remaining: %d, resetAt: %v ", allowed, remaining, resetAt)

		if !allowed {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "Rate limit exceeded"}`))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) GetRequestKey(path, ip string) string {
	return fmt.Sprintf("rate-limit:%s:%s", path, ip)
}

func (rl *RateLimiter) CheckRateLimit(
	ctx context.Context,
	key string,
) (allowed bool, remaining int, resetAt time.Time, err error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	fmt.Printf("Rate limiter check for: %s\n", key)

	now := time.Now()                          // e.g. 15:01:30
	windowStart := now.Add(-1 * rl.timeWindow) // 15:00:30

	// Use pipeline for efficiency
	pipe := rl.client.Pipeline()

	// Remove anything from sorted set (ZSET) before the windowStart time
	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart.UnixNano()))

	// Count number of entries (will be num of entries in window since we removed everything with earlier timestamp)
	zcard := pipe.ZCard(ctx, key)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return false, 0, time.Time{}, fmt.Errorf("redis pipeline error: %w", err)
	}

	numOfRequests := zcard.Val()

	if numOfRequests >= int64(rl.limit) {
		// Get oldest entry to calculate reset time
		oldestEntries, err := rl.client.ZRange(ctx, key, 0, 0).Result()
		if err != nil || len(oldestEntries) == 0 {
			resetAt = now.Add(rl.timeWindow)
		} else {
			// Parse timestamp from oldest entry
			oldestTime, err := strconv.ParseInt(oldestEntries[0], 10, 64)
			if err != nil {
				return false, 0, time.Time{}, fmt.Errorf("redis pipeline error: %w", err)
			}

			resetAt = time.Unix(0, oldestTime).Add(rl.timeWindow)
		}

		return false, 0, resetAt, nil
	}

	// Not rate limited
	// -------------------------

	// Add current request to sorted set
	// Score = timestamp (for sorting), Member = timestamp
	score := float64(now.UnixNano())
	member := fmt.Sprintf("%d", now.UnixNano())

	pipe2 := rl.client.Pipeline()
	pipe2.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	})

	// Set expiration to window duration to prevent memory leaks
	// Add buffer to handle edge cases
	pipe2.Expire(ctx, key, rl.timeWindow+(30*time.Second))

	_, err = pipe2.Exec(ctx)
	if err != nil {
		return false, 0, time.Time{}, fmt.Errorf("redis add error: %w", err)
	}

	remaining = rl.limit - int(numOfRequests) - 1 // -1 for the request we just added
	resetAt = now.Add(rl.timeWindow)

	return true, remaining, resetAt, nil
}

func getIPAddress(r *http.Request) string {
	// Check X-Forwarded-For header (if behind proxy/load balancer)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header (alternative header)
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Return as-is if parsing fails
	}

	return ip
}

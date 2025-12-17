package inmemory

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type RateLimiter struct {
	mu         sync.Mutex
	limit      int
	timeWindow time.Duration

	// Map structure: key -> list of request timestamps
	// key format: "path:ip" e.g. "/upload:192.168.1.1"
	requests map[string][]time.Time
}

type OptionFunc func(o *RateLimiter)

func WithLimit(limit int) OptionFunc {
	return func(o *RateLimiter) {
		o.limit = limit
	}
}

func New(opts ...OptionFunc) *RateLimiter {
	// defaults
	rl := &RateLimiter{
		limit:      10,
		timeWindow: 1 * time.Minute,
		requests:   make(map[string][]time.Time),
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

		allowed, remaining, resetAt := rl.CheckRateLimit(key)

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
	return fmt.Sprintf("%s:%s", path, ip)
}

func (rl *RateLimiter) CheckRateLimit(
	key string,
) (allowed bool, remaining int, resetAt time.Time) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	fmt.Printf("Rate limiter check for: %s requests %v\n", key, rl.requests[key])

	now := time.Now()                          // e.g. 15:01:30
	windowStart := now.Add(-1 * rl.timeWindow) // 15:00:30

	requests := rl.requests[key]

	// Get only the requests with timestamp after windowStart (e.g. in last minute)
	validRequests := make([]time.Time, 0)
	for _, req := range requests {
		if req.After(windowStart) {
			validRequests = append(validRequests, req)
		}
	}

	if len(validRequests) > rl.limit {
		// Rate limited
		// Find the oldest request to determine when the window resets
		oldestTimestamp := validRequests[0]
		resetAt = oldestTimestamp.Add(rl.timeWindow)

		return false, 0, resetAt
	}

	validRequests = append(validRequests, now)

	rl.requests[key] = validRequests
	remaining = rl.limit - len(validRequests)
	resetAt = now.Add(rl.timeWindow)

	return true, remaining, resetAt
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

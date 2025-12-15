package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JDGarner/go-playground/systems/ratelimiter"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("shutting down gracefully")
}

func run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	rl := ratelimiter.New(ratelimiter.WithLimit(3))

	handler := chain(
		mux,
		rl.Middleware,
		logMiddleware,
	)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type MiddlewareFunc func(next http.Handler) http.Handler

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func chain(handler http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging logs the request timestamp, method, and path.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := time.Now().UTC().Format(time.RFC3339)
		log.Printf("[%s] %s %s", ts, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// RequestTiming measures the time taken to serve a request, sets an X-Response-Time header,
// and logs the elapsed duration.
func RequestTiming(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		w.Header().Set("X-Response-Time", elapsed.String())
		log.Printf("Handled %s %s in %s", r.Method, r.URL.Path, elapsed)
	})
}

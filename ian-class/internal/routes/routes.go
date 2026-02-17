package routes

import (
	"net/http"

	"github.com/ianburns17/lab2-Ianburns/internal/handlers"
	"github.com/ianburns17/lab2-Ianburns/internal/middleware"
)

// SetupRoutes registers all routes on the provided mux and returns the configured mux.
func SetupRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/health", handlers.Health)

	// Register application routes and apply middleware (Logging -> RequestTiming).
	mux.Handle("/about", middleware.Logging(middleware.RequestTiming(http.HandlerFunc(handlers.About))))
	mux.Handle("/contact", middleware.Logging(middleware.RequestTiming(http.HandlerFunc(handlers.Contact))))
	mux.Handle("/qoute", middleware.Logging(middleware.RequestTiming(http.HandlerFunc(handlers.Qoute))))
	mux.Handle("/hello", middleware.Logging(middleware.RequestTiming(http.HandlerFunc(handlers.Hello))))
	// Root route
	mux.Handle("/", middleware.Logging(middleware.RequestTiming(http.HandlerFunc(handlers.Home))))

	return mux
}

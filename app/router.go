package app

import (
	chilogger "github.com/766b/chi-logger"
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/andyglass/go-restapi-template/app/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *Server) getRouter() *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(
		cors.Handler,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.Heartbeat("/ping"),
		chilogger.NewLogrusMiddleware("app", log),
		chiprometheus.NewMiddleware("go-restapi-template"),
	)

	r.Group(func(telemetry chi.Router) {
		telemetry.Get("/metrics", promhttp.Handler().ServeHTTP)
	})

	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/users", users.GetRouter())
	})

	return r
}

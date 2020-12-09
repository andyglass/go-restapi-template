package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/andyglass/go-restapi-template/app/config"
	"github.com/andyglass/go-restapi-template/app/logger"
)

var (
	log        = logger.Logger
	httpServer *http.Server
)

// Server global context
type Server struct {
	Cfg *config.Config
	DB  *sql.DB
}

// Run server
func (s *Server) Run() error {
	srv := &Server{
		Cfg: config.Get(),
		// DB:  db.Get(config.Cfg.DBConnStr),
	}
	config.Cfg = srv.Cfg
	// db.DB = srv.DB

	if err := logger.InitSentry(srv.Cfg.SentryDSN, srv.Cfg.Version); err != nil {
		log.Fatal(err)
	}

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", srv.Cfg.ServerHost, srv.Cfg.ServerPort),
		Handler:      s.getRouter(),
		ReadTimeout:  srv.Cfg.ReadTimeout,
		WriteTimeout: srv.Cfg.WriteTimeout,
	}
	log.Infof("Starting web server on port %s", srv.Cfg.ServerPort)

	return httpServer.ListenAndServe()
}

// Shutdown server
func (s *Server) Shutdown(ctx context.Context) error {
	return httpServer.Shutdown(ctx)
}

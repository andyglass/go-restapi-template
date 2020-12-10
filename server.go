package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/andyglass/go-restapi-template/config"
	"github.com/andyglass/go-restapi-template/logger"
)

var httpServer *http.Server

// Server global context
type Server struct {
	Cfg *config.Config
	DB  *sql.DB
}

// NewServer constructor
func NewServer() *Server {
	return &Server{
		Cfg: config.Get(),
		// DB:  db.Get(config.Cfg.DBConnStr),
	}
}

// Run server
func (s *Server) Run() error {
	config.Cfg = s.Cfg
	// db.DB = srv.DB

	if err := logger.InitSentry(s.Cfg.SentryDSN, s.Cfg.Version); err != nil {
		log.Fatal(err)
	}

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.Cfg.ServerHost, s.Cfg.ServerPort),
		Handler:      s.getRouter(),
		ReadTimeout:  s.Cfg.ReadTimeout,
		WriteTimeout: s.Cfg.WriteTimeout,
	}
	log.Infof("Starting web server on port %s", s.Cfg.ServerPort)

	return httpServer.ListenAndServe()
}

// Shutdown server
func (s *Server) Shutdown(ctx context.Context) error {
	return httpServer.Shutdown(ctx)
}

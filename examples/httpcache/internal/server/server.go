// Package server provides functionality to create and run an HTTP server.
package server

import (
	"httpcache/internal/server/respcache"
	"log/slog"
	"net/http"
)

// Server contains required information to
// run an HTTP server.
type Server struct {
	log   *slog.Logger
	cache *respcache.Cache
	serv  *http.Server
}

// NewServer creates a new Server instance with
// the specified address.
func NewServer(addr string) *Server { _ = "STUB: not implemented"; return nil }

// Start starts the server. It blocks until the server.Stop is called.
func (s *Server) Start() error { _ = "STUB: not implemented"; return nil }

// Stop shuts down the server.
func (s *Server) Stop() error { _ = "STUB: not implemented"; return nil }

// router sets up the HTTP routes for the server.
func (s Server) router() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// fetchReport is the handler that fetches report information based on
// the report name provided in the URL path.
func (s Server) fetchReport(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// For the demonstration purposes, the timer below acts
// as a placeholder for the actual report fetching logic.

// OK.

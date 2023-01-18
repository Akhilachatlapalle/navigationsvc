package service

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/betalo-sweden/navigationsvc/internal/navigationsvc/domain"
	"github.com/betalo-sweden/navigationsvc/internal/navigationsvc/transport"

	// third-party
	"github.com/go-chi/chi"
	"go.uber.org/zap"

	// local
	"github.com/betalo-sweden/navigationsvc/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(ctx context.Context, conf *config.Config) (*Server, error) {
	r := chi.NewRouter()
	r.Use(RequestLogger(conf.Logger))
	r.Use(ResponseLogger(conf.Logger))

	navigationService := domain.NewService(conf.Logger, conf.SectorID)
	navigationHandler := transport.NewHandler(conf.Logger, navigationService)

	r.Group(func(r chi.Router) {
		r.Post("/location", navigationHandler.GetLocation)
		r.Get("/", navigationHandler.Health)
	})

	httpServer := http.Server{
		Addr:    net.JoinHostPort("", conf.Port),
		Handler: r,
	}

	srv := Server{
		httpServer: &httpServer,
	}

	return &srv, nil
}

// Start starts the service.
func (s *Server) Start() error {
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop stops the service.
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

// URL returns the server URL.
func (s *Server) URL() string {
	return s.httpServer.Addr
}

// RequestLogger returns a middleware that logs incoming HTTP requests.
func RequestLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Incoming request", zap.String("url", r.RequestURI), zap.String("host", r.Host))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// ResponseLogger returns a middleware that logs outgoing HTTP responses.
func ResponseLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Info("Outgoing response",
				zap.Duration("latency", time.Since(start)),
			)

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

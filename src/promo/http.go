package promo

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

// HTTPServer holds http server info
type HTTPServer struct {
	binding      string
	httpListener *http.Server
	quit         chan os.Signal
	log          logrus.FieldLogger
	done         chan struct{}
}

// NewHTTPServer creates new http server
func NewHTTPServer(binding string, log logrus.FieldLogger) *HTTPServer {
	return &HTTPServer{
		binding: binding,
		log: log.WithFields(logrus.Fields{
			"prefix": "promo.http",
		}),
		quit: make(chan os.Signal, 1),
		done: make(chan struct{}),
	}
}

// Run starts http listener and returns error if any
func (s *HTTPServer) Run() error {
	log := s.log
	log.Infof("HTTP service start at %s", s.binding)
	defer log.Info("HTTP service stop")
	signal.Notify(s.quit, os.Interrupt)

	// r := s.setupRouter(SecureDecorator)

	s.httpListener = &http.Server{
		Addr:         s.binding,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		// Handler:      r,
	}

	errorC := make(chan error)
	go func() {
		if err := s.httpListener.ListenAndServe(); err != nil {
			// log.Error(err)
			errorC <- err
		}
	}()

	select {
	case err := <-errorC:
		return err
	case <-s.quit:
		return s.Shutdown()
	}
}

// Shutdown shuts down the http listener
func (s *HTTPServer) Shutdown() error {
	s.log.Info("HTTP service shutting down")
	close(s.done)

	// Create a deadline to wait for.
	wait := time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	return s.httpListener.Shutdown(ctx)
}

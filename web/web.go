package web

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type Web struct {
	config *Config
	log    *log.Entry
	router *mux.Router
}

func NewWeb(config *Config) *Web {
	a := &Web{
		log:    log.WithField("component", "api"),
		config: config,
		router: mux.NewRouter(),
	}
	a.registerMiddlewares()
	a.registerRoutes()
	return a
}

func (w *Web) registerMiddlewares() {
	w.router.Use(w.loggingMiddleware)
	w.router.Use(prometheusMiddleware)
	w.router.Use(jsonMiddleware)
}

func (w *Web) registerRoutes() {
	w.router.Path("/metrics").Handler(promhttp.Handler()).Methods("GET")
	w.router.Path("/").HandlerFunc(w.hello).Methods("GET")
}

func (w *Web) Serve() error {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	server := &http.Server{
		Addr:              w.config.Listen,
		Handler:           w.router,
		TLSConfig:         tlsConfig,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	w.log.Infof("Starting server on %s", server.Addr)
	return server.ListenAndServe()
}

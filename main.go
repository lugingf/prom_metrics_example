package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func main() {
	address := ":8080"
	server := http.Server{
		Addr:    address,
		Handler: NewRouter(),
	}

	InitMetrics()

	go runMetricServer()

	log.Printf("Listening on %s", address)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe failed")
	}
}

func runMetricServer() {
	mh := mux.NewRouter()
	mh.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

	listenAddress := ":9090"
	srv := &http.Server{
		Addr:        listenAddress,
		Handler:     mh,
		ReadTimeout: 1 * time.Second,
	}

	log.Printf(fmt.Sprintf("starting Metric exporter server: listening on %s", listenAddress))

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal("failed to listen promhandler server")
	}
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.Use(PrometheusMiddleware)

	r.HandleFunc("/", DefaultHandler)

	return r
}

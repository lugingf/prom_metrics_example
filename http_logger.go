package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status       int
	ResponseBody string
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *StatusRecorder) Write(body []byte) (int, error) {
	r.ResponseBody = string(body)
	return r.ResponseWriter.Write(body)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		route := mux.CurrentRoute(r)
		path, err := route.GetPathTemplate()
		if err != nil {
			log.Println("http_logger: get path template")
		}
		request, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("http_logger: read request body")
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(request))

		start := time.Now()

		next.ServeHTTP(recorder, r)

		SaveHTTPDurationHistogram(start, path, recorder.Status, r.Method)
		SaveHTTPDuration(start, path, recorder.Status, r.Method)
		SaveHTTPCount(1, path, recorder.Status, r.Method)

		sleepData := r.URL.Query().Get("sleep")
		sl, _ := strconv.Atoi(sleepData)
		SaveHTTPGauge(float64(sl), path, recorder.Status, r.Method)
	})
}

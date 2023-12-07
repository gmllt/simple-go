package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func (w *Web) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		jsonQuery, _ := json.Marshal(r.URL.Query())
		w.log.Println(fmt.Sprintf("[%s] %s \"%s\"", r.Method, r.RequestURI, jsonQuery))
		next.ServeHTTP(wr, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(wr, r)
	})
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(gHTTPDuration.WithLabelValues(path))
		rw := NewResponseWriter(wr)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode
		gResponseStatus.WithLabelValues(path, strconv.Itoa(statusCode)).Inc()
		gTotalRequests.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	})
}

func WriteJSON(wr http.ResponseWriter, data interface{}) {
	wr.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(wr).Encode(data)
}

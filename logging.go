package main

import (
	"github.com/fatih/color"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request information
		color.White("Request method: %s", r.Method)
		color.White("Request URL: %s", r.URL.String())

		// Record the status code using a custom response writer
		recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)

		// Log the response status code
		color.White("Response status: %d", recorder.statusCode)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

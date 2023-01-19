package middlewares

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		duration := time.Since(start).Seconds()
		statusCode := lrw.statusCode

		log.SetFormatter(&log.JSONFormatter{})
		if statusCode > http.StatusInternalServerError {
			log.WithFields(log.Fields{
				"duration": duration,
				"host":     req.Host,
				"path":     req.URL.Path,
				"method":   req.Method,
				"status":   statusCode,
			}).Error(http.StatusText(statusCode))
		} else if statusCode > http.StatusOK && statusCode < http.StatusInternalServerError {
			log.WithFields(log.Fields{
				"duration": duration,
				"host":     req.Host,
				"path":     req.URL.Path,
				"method":   req.Method,
				"status":   statusCode,
			}).Warn(http.StatusText(statusCode))
		} else {
			log.WithFields(log.Fields{
				"duration": duration,
				"host":     req.Host,
				"path":     req.URL.Path,
				"method":   req.Method,
				"status":   statusCode,
			}).Info(http.StatusText(statusCode))
		}
	})
}

package httptools

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type statusCodeWriter struct {
	http.ResponseWriter
	statuscode int
}

func (w *statusCodeWriter) WriteHeader(code int) {
	w.statuscode = code
	w.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusCodeWriter{
			ResponseWriter: w,
		}
		defer func() {
			log.WithContext(r.Context()).
				WithFields(log.Fields{
					"host":        r.Host,
					"method":      r.Method,
					"status_code": sw.statuscode,
					"time":        time.Since(start),
				}).Info(r.URL.RequestURI())
		}()
		next.ServeHTTP(sw, r)
	})
}

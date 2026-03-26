package router

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger := logrus.StandardLogger()
			start := time.Now()
			next.ServeHTTP(w, r)
			end := time.Now()
			logger.Infof("[ %s ] %s. Длительность выполнения:  %d мс.", r.Method, r.URL.Path, end.Sub(start).Milliseconds())
		})
}

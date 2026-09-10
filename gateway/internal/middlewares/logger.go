package middlewares

import (
	"context"
	"net/http"
	"time"
	"uuid"

	log "github.com/sirupsen/logrus"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func CustomLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestID := uuid.New()

		ctx := context.WithValue(
			r.Context(),
			requestIDKey,
			requestID,
		)

		log.Printf(
			"REQUEST id=%s method=%s path=%s remote=%s",
			requestID,
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r.WithContext(ctx))

		log.Printf(
			"REQUEST FINISHED id=%s method=%s path=%s duration=%s",
			requestID,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

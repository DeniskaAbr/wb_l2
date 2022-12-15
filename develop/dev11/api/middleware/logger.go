package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.HandlerFunc, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		logger.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

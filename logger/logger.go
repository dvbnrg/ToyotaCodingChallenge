package logger

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		pid := os.Getpid()

		log.Printf(
			"%s\t%s\t%s\t%s\t%v\t%s\t%s\t%v",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start).String(),
			pid,
			r.UserAgent(),
			r.Host,
			r.RemoteAddr,
			// w.Header,
			// r.Response.Body,
		)
	})
}

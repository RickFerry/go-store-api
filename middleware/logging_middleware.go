package middleware

import (
    "log"
    "net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware performs some action either before or after a request.

// This middleware function prints time and url path information to the terminal for each incoming request.
func LogRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%d/%d/%d : %d/%d ", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute())
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r) // move on to the next data that we want to serve
	})
}

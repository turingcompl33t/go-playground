package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RequestTimer(h http.Handler) http.Handler {
	// The generic signature for middleware is a function that
	// accepts an http.Handler and returns an http.Handler
	//
	// The http.Handler that is returned is (typically) a closure
	// that is converted to an http.HandlerFunc

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMessage = []byte("Incorrect password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMessage)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	const port = 8080

	securityProvider := TerribleSecurityProvider("GOPHER")

	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	}))
	wrappedMux := securityProvider(RequestTimer(mux))

	s := http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      wrappedMux,
	}
	fmt.Printf("Listening on localhost:%v\n", port)
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

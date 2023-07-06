package main

import (
	"fmt"
	"net/http"
	"time"
)

// A trivial implementation of the http.Handler interface
type HelloHandler struct{}

// ServeHTTP is the sole method defined on the http.Handler interface, always with this signature
func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func main() {
	const port = 8080

	s := http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      HelloHandler{},
	}
	fmt.Printf("Listening on localhost:%v\n", port)
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

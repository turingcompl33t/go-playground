package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const port = 8080

	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})

	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good dog!\n"))
	})

	// http.ServeMux implements http.Handler, so we can create a hierarchy of
	// request handlers that manage related collections of requests
	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))

	s := http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	fmt.Printf("Listening on localhost:%v\n", port)
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

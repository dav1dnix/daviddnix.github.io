package main

import (
	"log"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:    ":5000",
		Handler: nil,
	}

	http.Handle("/", http.FileServer(http.Dir("./web")))
	log.Println("Listening on port: ", s.Addr)
	s.ListenAndServe()
}

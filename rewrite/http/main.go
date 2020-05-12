package main

import (
	"net/http"
	"log"
)

func main() {
	s := &http.Server{
		Addr: ":5000",
		Handler: nil
	}

	http.Handle("/", http.FileServer(http.Dir(web)))
	s.ListenAndServe()
}

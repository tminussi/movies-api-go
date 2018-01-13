package main

import (
	"net/http"
	"github.com/tminussi/hello/movie/handler"
)

func main() {
	http.HandleFunc("/api/movies", handler.MovieHandler)
	http.ListenAndServe(":8080", nil)
}
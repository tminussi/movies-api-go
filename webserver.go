package main

import (
	"net/http"
	"github.com/tminussi/movie-api/movie/handler"
)

func main() {
	http.HandleFunc("/api/movies", handler.MovieHandler)
	http.ListenAndServe(":8080", nil)
}
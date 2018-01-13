package handler

import (
	"net/http"
	"encoding/json"
	"github.com/tminussi/movie-api/movie/model"
	"github.com/tminussi/movie-api/movie/repository"
)

func MovieHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMovies(w)
	case "POST":
		createMovie(w, r)
	}

}

func getMovies(w http.ResponseWriter) {
	err := repository.OpenConnection()
	if err != nil {
		panic(err)
	}
	movies := repository.GetMovies()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	err := repository.OpenConnection()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(r.Body)
	newMovie := movie.Movie{}
	decoder.Decode(&newMovie)
	id := repository.InsertMovie(newMovie)
	newMovie.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

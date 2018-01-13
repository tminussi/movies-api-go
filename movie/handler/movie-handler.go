package handler

import (
	"net/http"
	"github.com/tminussi/hello/movie/model"
	"encoding/json"
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
	movieOne := movie.Movie{
		Id:1,
		Name:"Devil's Advocate",
		Genre: "Mystery",
	}
	movieTwo := movie.Movie{
		Id:2,
		Name:"The Godfather",
		Genre: "Ganster",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	movies := []movie.Movie{}
	movies = append(movies, movieOne)
	movies = append(movies, movieTwo)
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	newMovie := movie.Movie{}
	decoder.Decode(&newMovie)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

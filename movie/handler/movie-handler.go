package handler

import (
	"net/http"
	"encoding/json"
	"github.com/tminussi/movie-api/movie/model"
	"github.com/tminussi/movie-api/movie/repository"
	"fmt"
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
	sql := "select id, name, genre from movies"
	if err != nil {
		fmt.Println("trying to connect to db", err)
	}
	movies := []movie.Movie{}
	err = repository.Db.Select(&movies, sql)
	if err != nil {
		panic(err)
	}
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
	sql := "INSERT INTO movies (NAME, GENRE) VALUES (?, ?)"
	result, _ := repository.Db.Exec(sql, newMovie.Name, newMovie.Genre)
	id, _ := result.LastInsertId()
	newMovie.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

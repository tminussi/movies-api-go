package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tminussi/movie-api/movie/model"
)

var Db *sqlx.DB

func OpenConnection() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/movies")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

func InsertMovie(movie movie.Movie) int64 {
	sql := "INSERT INTO movies (NAME, GENRE) VALUES (?, ?)"
	result, _ := Db.Exec(sql, movie.Name, movie.Genre)
	id, _ := result.LastInsertId()
	return id
}

func GetMovies() []movie.Movie {
	sql := "select id, name, genre from movies"
	movies := []movie.Movie{}
	err := Db.Select(&movies, sql)
	if err != nil {
		panic(err)
	}
	return movies
}

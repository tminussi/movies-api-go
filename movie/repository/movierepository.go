package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
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

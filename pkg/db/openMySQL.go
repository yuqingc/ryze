package db

import (
	"database/sql"
)

// OpenMySQL no description
func OpenMySQL() (db *(sql.DB), err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ryze")
	return
}

package db

import (
	"database/sql"
	"fmt"
)

// OpenMySQL no description
func OpenMySQL() (db *(sql.DB), err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ryze")
	if err != nil {
		err = fmt.Errorf("fail to open mySQL: %v", err)
	}
	return
}

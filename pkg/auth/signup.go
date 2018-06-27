package auth

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// WriteUserToDB need a mysql server running in the system first
func WriteUserToDB(name, password string) (string, error) {
	fmt.Println("Start writing user to database")
	cryptedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return `brypto failed`, err
	}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ryze")
	if err != nil {
		fmt.Println(err)
		return `open db failed`, err
	}
	defer db.Close()

	userIns, err := db.Prepare("insert into passwd SET username=?, password=?")
	if err != nil {
		fmt.Println(err)
		return `prepare failed`, err
	}
	defer userIns.Close()

	_, err = userIns.Exec(name, string(cryptedPasswd))
	if err != nil {
		fmt.Println(err)
		return `exec failed`, err
	}
	return name + ` written`, nil
}

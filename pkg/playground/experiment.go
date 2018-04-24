package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// docker run --name mt-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:5.7
// docker run --name mt-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
// docker exec -it mt-mysql bin/bash

func main() {
	fmt.Println("playground")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ryze")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	// userIns, err := db.Prepare("insert passwd SET username=?, password=?")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// defer userIns.Close()
	// res, err := userIns.Exec("admin", "$2a$10$WvY1m3O6MMET2mwLjSvQXOGYE7gYGxyaQFf53Ds1t7mEFEQwEwK.W")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// fmt.Println(res)
	rows, err := db.Query("select * from passwd")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rows.Next() {
		var name string
		var password string
		err = rows.Scan(&name, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, password)
	}
}

package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuqingc/ryze/pkg/auth"
)

var name = "yuqingc"

var password = "bdmin"

func main() {
	res, err := auth.WriteUserToDB(name, password)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}

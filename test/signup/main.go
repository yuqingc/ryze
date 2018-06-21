package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuqingc/ryze/pkg/auth"
)

func main() {
	res, err := auth.WriteUserToDB("admin", "bdmin")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuqingc/ryze/pkg/auth"
)

func main() {
	writeRes, err := auth.WriteUserToDB("admin", "pwd123456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(writeRes)
}

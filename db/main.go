package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("failed to connect to the database: %s", err)
	}
	fmt.Println("hello world")
	defer db.Close()
}

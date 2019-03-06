package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	)

//Initialize the database

//为何要申明变量 --> 因为想定义一个dbops包下的局部变量
var (
	dbConn *sql.DB
	err error
)

//sql.Open是go的自带数据库操作库。其中sql是package名，Open是这个包中的一个函数。
//open函数返回两个值，类型为*DB, error。 *DB是一个叫做DB的结构体的指针，error是一个结构体，结构体中包含一个Error（）的函数。
func init() {
	dbConn, err = sql.Open("mysql", "root:11111111@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())   //为何err.Error()才能panic？？？
	}
}
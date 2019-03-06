package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//注册处理程序
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}


func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

//notes
//listen -> RegisterHandlers -> Handlers

//the detail of the handler
//handler -> validation{1.request, 2.user} -> business logic -> response

//validation
//1. data model
//2. error

//business logic (at dbops)
//additions, deletions, changes, search of database

//session
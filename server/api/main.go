package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//定义一个结构体 这个结构体的r是Router结构体
type middleWareHandler struct{
	r *httprouter.Router
}

//新建一个middleWareHandler
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}


//注册处理程序
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}


func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
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


//main -> middleware -> defs(message, err) -> handelers -> dbops -> response
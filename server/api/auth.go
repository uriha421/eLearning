package main

import (
	"github.com/eLearning/server/api/defs"
	"net/http"
	"github.com/eLearning/server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

//session校验
func validateUserSession(r *http.Request) bool {
	//判断是否有session 没有就返回false
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	//判断session是否过期 如果过期就返回false 没有过期就返回true
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	//？？为什么要把uname加进去？？
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true

}

//user校验
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_SESSION)
	if len(uname) == 0 {
		SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}
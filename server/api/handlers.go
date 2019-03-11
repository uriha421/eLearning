package main

import (
	"github.com/eLearning/server/api/defs"
	"github.com/eLearning/server/api/session"
	"encoding/json"
	"github.com/eLearning/server/api/dbops"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//io.WriteString(w, "Create User Handler")

	//读取请求的内容
	res, _ := ioutil.ReadAll(r.Body)
	//定义一个UserCredential类型的变量 并将其指针赋值给ubody
	ubody := &defs.UserCredential{}

	//将res（json）反序列化成ubody（结构体）
	if err := json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	//将从ubody写入数据库（通过AddUserCredential）
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil{
		SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	//生成一个session
	id := session.GenerateNewSessionId(ubody.Username)
	//将signup类型的结构体写入su
	su := &defs.SignedUp{Success:true, SessionId: id }

	//将su序列化并写入resp
	if resp, err := json.Marshal(su); err != nil{
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		SendNormalResponse(w, string(resp), 201)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/eLearning/server/api/defs"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse){
	//把定义好的HttpSC直接写到head里面
	w.WriteHeader(errResp.HttpSC)

	//序列化成json string
	resStr, _ := json.Marshal(&errResp.Error)

	//通过io将数据写出去
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int){
	//将sc写入header中
	w.WriteHeader(sc)
	//将w和resp写出去
	io.WriteString(w, resp)
}
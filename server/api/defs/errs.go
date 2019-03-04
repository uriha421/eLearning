package defs

//the definition of error

type Err struct {
	Error string       `json:"error"`
	ErrorCode string   `json:"error_code"`
}

type ErroResponse struct {
	HttpSC int
	Error Err
}

var (
	//错误请求 正文解析失败
	ErrorRequestBodyParseFailed = ErroResponse{HttpSC:400, Error: Err{Error:"Requset body is not correctly", ErrorCode:"001"}}
	//用户验证错误
	ErrorNotAuthUser = ErroResponse{HttpSC:400,Error:Err{Error:"User authentication failed", ErrorCode:"002"}}
)

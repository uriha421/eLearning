package defs

//the definition of error

type Err struct {
	Error string       `json:"error"`
	ErrorCode string   `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	//错误请求 正文解析失败
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC:400, Error: Err{Error:"Requset body is not correctly", ErrorCode:"001"}}
	//用户验证错误
	ErrorNotAuthUser = ErrResponse{HttpSC:401,Error:Err{Error:"User authentication failed", ErrorCode:"002"}}
	//数据库操作错误（内部错误）
	ErrorDBError = ErrResponse{HttpSC:500,Error:Err{Error:"DB ops failed", ErrorCode:"003"}}
	//内部错误
	ErrorInternalFaults = ErrResponse{HttpSC:500, Error:Err{Error:"Internal service error", ErrorCode:"004" }}
)

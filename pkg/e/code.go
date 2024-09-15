package e

// 自定义状态码，方便api结果处理
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400 //无效参数

	//token相关的错误处理
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

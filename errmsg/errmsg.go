//errMsg包
package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code=1000... 用户模块的错误
	ErrorUsernameUsed   = 1001 //用户名错误（用户名已经被使用了。）
	ErrorPasswordWrong  = 1002 //用户密码错误。
	ErrorUserNotExist   = 1003 //用户不存在。
	ErrorTokenExist     = 1004 //用户携带token不存在。
	ErrorTokenRuntime   = 1005 //token超时过期。
	ErrorTokenWrong     = 1006 //用户携带的token和我们提供的不一致。
	ErrorTokenTypeWrong = 1007 //token格式错误(token是有格式要求的。)

	//code=2000... 文章模块的错误

	//code=3000... 分类模块的错误
)

//字典
//string是我们要抛出的错误信息
//int就是code状态码
var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "error",
	ErrorUsernameUsed:   "用户名已存在!",
	ErrorPasswordWrong:  "用户密码错误!",
	ErrorUserNotExist:   "用户不存在!",
	ErrorTokenExist:     "TOKEN不存在!",
	ErrorTokenRuntime:   "TOKEN超时过期!",
	ErrorTokenWrong:     "TOKEN不正确!",
	ErrorTokenTypeWrong: "TOKEN格式错误!",
}

//输出错误信息的函数
func GetErrMsg(code int) string {
	return codeMsg[code]
}

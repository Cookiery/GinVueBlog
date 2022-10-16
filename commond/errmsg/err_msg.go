package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004 // TOKEN不存在
	ERROR_TOKEN_RUNTIME    = 1005 // TOKEN超时过期
	ERROR_TOKEN_WRONG      = 1006 // TOKEN错误（虚假TOKEN）
	ERROR_TOKEN_TYPE_WRONG = 1007 // TOKEN格式错误

	// 2000... 分类模块的错误
	ERROR_CATEGORY_USED = 2001 // 分类已经存在

	// 3000... 文章模块的错误

)

var errorMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户名不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",

	ERROR_CATEGORY_USED: "分类已经存在",
}

func ErrMsg(code int) string {
	return errorMsg[code]
}

package httpx

var message map[ErrorType]string

type ErrorType uint32

// 成功返回
const OK ErrorType = 200

/**(前3位代表业务,后三位代表具体功能)**/
//全局错误码
const (
	SERVER_COMMON_ERROR ErrorType = iota + 100001
	REUQEST_PARAM_ERROR
	TOKEN_EXPIRE_ERROR
	TOKEN_GENERATE_ERROR
	DB_ERROR
	MOBILE_NOTFOUND
	PASSWORD_ERR
	QUERY_ERR
	FILE_UPLOAD_ERR
	WXMINI_AUTH_ERR
	API_LIMIT_ERR
	BACKLIST_ERR
	RESUBMIT_ERR
)

func init() {
	message = make(map[ErrorType]string)
	message[OK] = "OK"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效,请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[MOBILE_NOTFOUND] = "手机号码不存在"
	message[PASSWORD_ERR] = "密码错误"
	message[QUERY_ERR] = "用户信息查询失败"
	message[FILE_UPLOAD_ERR] = "文件上传失败"
	message[WXMINI_AUTH_ERR] = "微信授权失败"
	message[API_LIMIT_ERR] = "接口限流"
	message[BACKLIST_ERR] = "拉入黑名单,请联系管理员处理."
	message[RESUBMIT_ERR] = "重复提交"
}

package serializer

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

const (

	//Token保存时间（天）
	TokenRedisTime = 5

	SystemOk = 0

	SystemError = 1

	// CodeCheckLogin 未登录
	CodeCheckLogin = 401

	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403

	// CodeDBError 数据库操作失败
	CodeDBError = 5000

	// 授权失败
	CodeAuthorizationFailed = 50001
	CodeTokenIllegalToken = 50008
	CodeTokenNull = 50005

	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)
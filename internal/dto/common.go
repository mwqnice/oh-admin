package dto

type SuccessResponse struct {
	Code int         `json:"code" example:"0"`      // 响应编码：0成功 401请登录 403无权限 500错误
	Msg  string      `json:"msg" example:"success"` // 消息提示
	Data interface{} `json:"data"`                  // 数据对象
}

type CaptchaResponse struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

package goxy

import "github.com/kataras/iris/v12/context"

const (
	StatusOK               = 0     //成功
	StatusFail             = 40000 //失败
	StatusParameterError   = 40001 //参数错误
	StatusDataNotExist     = 40002 //数据不存在
	StatusValidationFailed = 40003 //验证失败
	StatusServerReason     = 40004 //服务器原因
)

type IrisHttpResult struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (h *IrisHttpResult) Error(c context.Context, code int64, msg string) {
	var resp = IrisHttpResult{
		Code: code,
		Msg:  msg,
	}
	c.JSONP(resp)
}

func (h *IrisHttpResult) Echo(c context.Context, code int64, data interface{}) {
	var resp = IrisHttpResult{
		Code: code,
		Msg:  "成功",
		Data: data,
	}
	c.JSONP(resp)
}

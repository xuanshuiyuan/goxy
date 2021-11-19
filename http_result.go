package goxy

import "github.com/kataras/iris/v12/context"

type HttpResult struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (h *HttpResult) Error(c context.Context, code int64, msg string) {
	var resp = HttpResult{
		Code: code,
		Msg:  msg,
	}
	c.JSONP(resp)
}

func (h *HttpResult) Echo(c context.Context, code int64, data interface{}) {
	var resp = HttpResult{
		Code: code,
		Msg:  "成功",
		Data: data,
	}
	c.JSONP(resp)
}

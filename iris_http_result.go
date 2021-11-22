package goxy

import "github.com/kataras/iris/v12/context"

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

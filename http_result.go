package goxy

import "github.com/kataras/iris/v12/context"

type Result struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Errors(c context.Context, code int64, msg string) {
	var resp = Result{
		Code: code,
		Msg:  msg,
	}
	c.JSONP(resp)
}

func Echo(c context.Context, code int64, data interface{}) {
	var resp = Result{
		Code: code,
		Msg:  "成功",
		Data: data,
	}
	c.JSONP(resp)
}

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

var sub_code = map[int64]string{
	StatusOK:               "Success",
	StatusFail:             "Failure",
	StatusParameterError:   "Parameter_Error",
	StatusDataNotExist:     "DataNot_Exist",
	StatusValidationFailed: "Validation_Failed",
	StatusServerReason:     "Server_Reason",
}

var sub_msg = map[int64]string{
	StatusOK:               "成功",
	StatusFail:             "失败",
	StatusParameterError:   "参数错误",
	StatusDataNotExist:     "数据不存在",
	StatusValidationFailed: "验证失败",
	StatusServerReason:     "服务器原因",
}

type IrisHttpResult struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	SubCode string      `json:"sub_code"`
	SubMsg  string      `json:"sub_msg"`
}

func (h *IrisHttpResult) Error(c context.Context, code int64, msg string) {
	var resp = IrisHttpResult{
		Code:    code,
		Msg:     msg,
		SubCode: sub_code[code],
		SubMsg:  sub_msg[code],
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

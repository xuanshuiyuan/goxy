package goxy

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/context"
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func LogInit() {
	//日志输出文件
	var path = "target/log/logs.log"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}
	//自定义日志格式
	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func ErrorFmt(c context.Context, msg string, data interface{}) {
	type LogType struct {
		method string
		path   string
		msg    string
		data   interface{}
		params interface{}
	}
	var logType = LogType{
		method: c.Method(),
		path:   c.Path(),
		params: c.FormValues(),
		data:   data,
		msg:    msg,
	}
	logStr := "\nMethod:%s \nPath:%s \nParams:%s \nMsg:%s \nData:%s"
	paramsStr, _ := json.Marshal(logType.params)
	dataStr, _ := json.Marshal(logType.data)
	Error.Println(fmt.Sprintf(logStr, logType.method, logType.path, paramsStr, msg, dataStr))
}

package goxy

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Logs struct {
}

//
//type Data struct {
//}

var F *os.File
var defalut = "defalut"
var objName *string = &defalut
var root, _ = os.Getwd()
var path = fmt.Sprintf("%s%s", root, "/target/logs/")
var file = "logs.log"
var year = time.Now().Year()
var month = time.Now().Month()
var day = time.Now().Day()
var lastname *string = &file

func (l *Logs) checkDir() (errs error) {
	var paths = fmt.Sprintf("%s%s/%d/%d%d/", path, *objName, year, month, day)
	_, err := os.Stat(paths) // 通过获取文件信息进行判断
	if err != nil {
		// 错误不为空，表示目录不存在
		err := os.MkdirAll(paths, 0755)
		//defer f.Close()
		if err != nil {
			// 创建文件失败处理
			return err
		}
	} else {
		// 错误为空，表示文件存在
		return nil
	}
	return nil
}

func (l *Logs) checkFile(filename string) (errs error) {
	_, err := os.Stat(filename) // 通过获取文件信息进行判断
	if err != nil {
		// 错误不为空，表示文件不存在
		f, err := os.Create(filename)
		defer f.Close()
		if err != nil {
			// 创建文件失败处理
			return err
		}
	} else {
		// 错误为空，表示文件存在
		return nil
	}
	return nil
}

func (l *Logs) OpenFile(filename string) {
	if err := l.checkDir(); err != nil {
		log.Fatalln("Faild to MkdirAllFile error logger file:", err)
	}
	if filename == "" {
		filename = file
	}
	var files = fmt.Sprintf("%s%s/%d/%d%d/%s", path, *objName, year, month, day, filename)
	if err := l.checkFile(files); err != nil {
		log.Fatalln("Faild to CreateFile error logger file:", err)
	}
	////日志输出文件
	f, err := os.OpenFile(files, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}
	F = f
}

func (l *Logs) before(obj string, filename string) {
	if filename == "" {
		filename = file
	}
	if obj == "" {
		obj = "default"
	}
	//fmt.Println(obj)
	if *lastname != filename || *objName != obj || F == nil {
		lastname = &filename
		objName = &obj
		l.OpenFile(filename)
	}
}

func (l *Logs) Info(obj string, filename string) *log.Logger {
	l.before(obj, filename)
	return log.New(io.MultiWriter(F, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *Logs) Error(obj string, filename string) *log.Logger {
	l.before(obj, filename)
	return log.New(io.MultiWriter(F, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *Logs) Data(obj string, filename string) *log.Logger {
	l.before(obj, filename)
	return log.New(io.MultiWriter(F, os.Stderr), "DATA: ", log.Ldate|log.Ltime|log.Lshortfile)
}

//func (d *Data) Add(v ...interface{}) *log.Logger {
//	for k, val := range v {
//		vs, _ := json.Marshal(val)
//		c := fmt.Sprintf("%s", vs)
//		v[k] = c
//	}
//	return log.New(io.MultiWriter(F, os.Stderr), "DATA: ", log.Ldate|log.Ltime|log.Lshortfile)
//}

package goxy

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Logs struct {
	Path string
}

var F *os.File
var ObjName string

var path = "target/logs/"
var file = "logs.log"
var year = time.Now().Year()
var month = time.Now().Month()
var day = time.Now().Day()
var lastname *string = &file

func (l *Logs) checkDir() (errs error) {
	var paths = fmt.Sprintf("%s%s/%d/%d%d/", l.Path, ObjName, year, month, day)
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
	var files = fmt.Sprintf("%s%s/%d/%d%d/%s", l.Path, ObjName, year, month, day, filename)
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

func (l *Logs) before(filename string) {
	if filename == "" {
		filename = file
	}
	rootPath, err := os.Getwd()
	if err != nil {
		ObjName = "obj"
	}
	wd := strings.Split(string(rootPath), "/")
	ObjName = wd[len(wd)-1]
	//fmt.Println(F)
	if *lastname != filename || F == nil {
		lastname = &filename
		l.OpenFile(filename)
	}
}

func (l *Logs) Info(filename string) *log.Logger {
	l.before(filename)
	return log.New(io.MultiWriter(F, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *Logs) Error(filename string) *log.Logger {
	l.before(filename)
	return log.New(io.MultiWriter(F, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// @Author xuanshuiyuan 2022/1/1 10:40:00
package goxy

import (
	"bytes"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"
)

func mapKeyExist(mapValue map[string]interface{}, key string) (bool, interface{}) {
	if value, ok := mapValue[key]; ok {
		return true, value
	} else {
		return false, nil
	}
}

func StringToInterface(string []string) (list []interface{}) {
	if reflect.TypeOf(string).Kind() == reflect.Slice {
		val := reflect.ValueOf(string)
		for i := 0; i < val.Len(); i++ {
			ele := val.Index(i)
			list = append(list, ele.Interface())
		}
	}
	return
}

// @Title GetJsonField
// @Description 获取json字符串中指定字段内容  ioutil.ReadFile()读取字节切片
// @Author xuanshuiyuan 2022-05-31 10:00
// @Param
// @Return
func GetJsonField(bytes []byte, field ...string) []byte {
	if len(field) < 1 {
		return nil
	}
	//将字节切片映射到指定map上  key：string类型，value：interface{}  类型能存任何数据类型
	var mapObj map[string]interface{}
	json.Unmarshal(bytes, &mapObj)
	var tmpObj interface{}
	tmpObj = mapObj
	for i := 0; i < len(field); i++ {
		tmpObj = tmpObj.(map[string]interface{})[field[i]]
		if tmpObj == nil {
			return nil
		}
	}
	result, err := json.Marshal(tmpObj)
	if err != nil {
		return nil
	}
	return result
}

// @Title StringSortCompare
// @Description 比较数组排序后是否一样
// @Author xuanshuiyuan 2022-03-01 14:07
// @Param
// @Return error
func StringSortCompare(arr1 []string, arr2 []string) (err error) {
	if len(arr1) == 0 || len(arr2) == 0 {
		return errors.New("array is empty")
	}
	sort.Strings(arr1)
	sort.Strings(arr2)
	if strings.Join(arr1, ",") != strings.Join(arr2, ",") {
		return errors.New("array unequal")
	}
	return
}

// @Title RemoveRepeatedElement
// @Description 数组去重
// @Author xuanshuiyuan 2022-01-24 17:18
// @Param string
// @Return string
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// @Title StringsInSearch
// @Description 查看数组是否存在某个字符串
// @Author xuanshuiyuan 2022-01-24 17:18
// @Param string
// @Return string
func StringsInSearch(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

// @Title StructAssign
// @Description 复制结构体
// @Author xuanshuiyuan 2022-01-06 10:25
// @Param string 2006-01-02
// @Return string
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

// @Title TimeParseInLocation
// @Description 字符串日期格式化成时间戳
// @Author xuanshuiyuan 2022-01-06 10:25
// @Param string 2006-01-02
// @Return string
func TimeParseInLocation(layout string, t string) int64 {
	t2, _ := time.ParseInLocation(layout, t, time.Local)
	return t2.Unix()
}

// @Title GetTimeForm
// @Description 格式化时间戳
// @Author xuanshuiyuan 2022-02-10 14:30
// @Param string 2006-01-02
// @Return string
func GetTimeForm(strTime int64) string {
	//记12345,3那个位置的数这里我使用的15，也就是用24小时格式来显示，如果直接写03则是12小时am pm格式。
	timeLayout := "2006-01-02 15:04:05"
	datetime := time.Unix(strTime, 0).Format(timeLayout)
	return datetime
}

// @Title GetTimeYmdForm
// @Description 格式化时间戳
// @Author xuanshuiyuan 2022-02-10 14:30
// @Param string 2006-01-02
// @Return string
func GetTimeYmdForm(strTime int64) string {
	//记12345,3那个位置的数这里我使用的15，也就是用24小时格式来显示，如果直接写03则是12小时am pm格式。
	timeLayout := "2006-01-02"
	datetime := time.Unix(strTime, 0).Format(timeLayout)
	return datetime
}

// @Title GetTimeYmForm
// @Description 格式化时间戳
// @Author xuanshuiyuan 2022-02-10 14:30
// @Param string 2006-01-02
// @Return string
func GetTimeYmForm(strTime int64) string {
	//记12345,3那个位置的数这里我使用的15，也就是用24小时格式来显示，如果直接写03则是12小时am pm格式。
	timeLayout := "2006-01"
	datetime := time.Unix(strTime, 0).Format(timeLayout)
	return datetime
}

// @Title TimeParse
// @Description 检测时间字符串是否正确
// @Author xuanshuiyuan 2022-01-06 10:25
// @Param string 2006-01-02
// @Return string
func TimeParse(layout string, t string) (err error) {
	_, err = time.Parse(layout, t)
	if err != nil {
		return
	}
	return
}

// @Title FormatTimeString
// @Description 格式化日期字符串
// @Author xuanshuiyuan 2022-01-06 10:25
// @Param string
// @Return string
func FormatTimeString(t string) string {
	var ret = ""
	timestr := strings.ReplaceAll(t, "/", "-")
	arr := strings.Split(timestr, " ")
	if len(arr) == 1 || len(arr) == 0 {
		ret = strings.Join([]string{arr[0], "00:00:00"}, " ")
	} else {
		switch strings.Count(arr[1], ":") {
		case 0:
			ret = strings.Join([]string{arr[0], strings.Join([]string{arr[1], ":00:00"}, "")}, " ")
			break
		case 1:
			ret = strings.Join([]string{arr[0], strings.Join([]string{arr[1], ":00"}, "")}, " ")
			break
		default:
			ret = timestr
			break
		}
	}
	return ret
}

// @Title Contain
// @Description 检测Slice,Array,Map是否包含某个值
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param obj,target
// @Return bool
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

// @Title CheckDir
// @Description 检测目录是否创建
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param path
// @Return error
func CheckDir(path string) error {
	_, err := os.Stat(path) // 通过获取文件信息进行判断
	if err != nil {
		// 错误不为空，表示目录不存在
		errs := os.MkdirAll(path, 0755)
		//defer f.Close()
		if errs != nil {
			// 创建文件失败处理
			return errs
		}
	} else {
		// 错误为空，表示文件存在
		return nil
	}
	return nil
}

// @Title StaticFileDirectory
// @Description 静态文件目录
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param
// @Return string
func StaticFileDirectory() string {
	var root, _ = os.Getwd()
	var path = fmt.Sprintf("%s%s", root, "/target/static")
	return fmt.Sprintf("%s/%s", path, YmdStr())
}

// @Title YmdStr
// @Description 当前年月日格式的字符串
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param
// @Return string
func YmdStr() string {
	var year = time.Now().Year()
	var month = time.Now().Format("01")
	var day = time.Now().Format("02")
	return fmt.Sprintf("%d/%s/%s", year, month, day)
}

// @Title Md5Str
// @Description 字符串md5加密
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param str
// @Return string
func Md5Str(str string) string {
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5String
}

// @Title StrVerify
// @Description 将最前面和最后面的ASCII定义的空格去掉
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param str
// @Return string
func StrVerify(str string) string {
	return strings.TrimSpace(str)
}

// @Title RedisTokenValue
// @Description redis token
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param
// @Return string
func RedisTokenValue() string {
	return RandChar(20)
}

// @Title RandChar
// @Description 生成随机字符串
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param size
// @Return string
func RandChar(size int) string {
	b := make([]byte, size)
	//ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return SeedRandChar(size)
	}
	return base64.URLEncoding.EncodeToString(b) //将生成的随机数b编码后返回字符串,该值则作为session ID
}

// @Title SeedRandChar
// @Description 生成随机字符串,非并发安全 seed
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param size
// @Return string
func SeedRandChar(size int) string {
	var char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var b bytes.Buffer
	for i := 0; i < size; i++ {
		b.WriteByte(char[rand.Int63()%int64(len(char))])
	}
	return b.String()
}

// @Title FmtLog
// @Description 格式化日志
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param v
// @Return interface{}
func FmtLog(v ...interface{}) interface{} {
	for k, val := range v {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Slice:
			vc := val.([]byte)
			v[k] = FmtJson(&vc)
		case reflect.Map, reflect.Struct, reflect.Ptr:
			vs, _ := json.Marshal(val)
			c := fmt.Sprintf("%s", vs)
			vc := []byte(c)
			v[k] = FmtJson(&vc)
		case reflect.String:
			if strings.Contains(val.(string), ".title") {
				val = val.(string)[:(len(val.(string)) - 6)]
				var buf bytes.Buffer
				buf.WriteString("\n")
				buf.WriteString("-----------")
				buf.WriteString(fmt.Sprintf("%s", val))
				buf.WriteString("-----------")
				buf.WriteString("\n")
				v[k] = buf.String()
			}

		}
	}
	return v
}

// @Title FmtJson
// @Description 格式化输出[]byte
// @Author xuanshuiyuan 2021-10-22 17:14:47
// @Param v
// @Return string
func FmtJson(v *[]byte) string {
	var str bytes.Buffer
	_ = json.Indent(&str, *v, "", "    ")
	return str.String()
}

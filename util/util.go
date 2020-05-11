package util

import (
	"cache-system/constant"
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/**
 获取日志等级的字符串
 */
func GetLevelText(level int)string{

	switch level {
	case constant.LogLevelDebug:
		return "DEBUG"
	case constant.LogLevelTrace:
		return "TRACE"
	case constant.LogLevelInfo:
		return "INFO"
	case constant.LogLevelWarn:
		return "WARN"
	case constant.LogLevelError:
		return "ERROR"
	case constant.LogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func GetLogLevel(level string) int{
	switch level {
	case "debug":
		return constant.LogLevelDebug
	case "trace":
		return constant.LogLevelTrace
	case "info":
		return constant.LogLevelInfo
	case "warn":
		return constant.LogLevelWarn
	case "error":
		return constant.LogLevelError
	case "fatal":
		return constant.LogLevelFatal
	default:
		return constant.LogLevelDebug
	}
}


type LogData struct {
	Message string
	TimeStr string
	LevelStr string
	Filename string
	FuncName string
	LineNo int
	WarnAndFatal bool
}

//获取执行文件名，调用函数名及当前执行位置
func GetLineInfo()(fileName string,funcName string,lineNo int){

	pc,file,line,ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}


/**
  1.当调用打印日志的方法时，我们把日志相关的数据写入到chan(队列)
  2.然后我们就有一个后台的线程不断从chan里读取这些日志，最终写入文件里
 */
func WriteLog(level int,format string,args...interface{})*LogData{
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := GetLevelText(level)

	fileName,funcName,lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format,args...)

	logData := &LogData{
		Message:      msg,
		TimeStr:      nowStr,
		LevelStr:     levelStr,
		Filename:     fileName,
		FuncName:     funcName,
		LineNo:       lineNo,
		WarnAndFatal: false,
	}

	if level == constant.LogLevelError || level == constant.LogLevelWarn || level == constant.LogLevelFatal{
		logData.WarnAndFatal = true
	}
	return logData
}

//1KB，100KB，1MB，2MB，1GB
func ParseSize(size string) int64{

	size = strings.TrimSpace(size)
	size = strings.ToLower(size)

	if size == ""{
		size = "256MB"
	}
	reg := regexp.MustCompile(`(\d+)(kb|mb|gb)`)
	find := reg.FindStringSubmatch(size)

	s := find[1]
	u := find[2]

	if s == "" {
		s = "256"
	}

	if u == "" {
		u = "mb"
	}


	rs,err := strconv.ParseInt(s,10,64)
	if err != nil {
		rs = 256
	}

	var r int64

	switch u {
	case "kb":
		r = rs * 1024
	case "mb":
		r = rs * 1024*1024
	case "gb":
		r = rs * 1024 * 1024 * 1024
	default:
		r = 256*1024*1024
	}
	return  r
}

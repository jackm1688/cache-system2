package logger

import "fmt"

var log ILogger //定义全局变量

/**
* file, 初始化一个日志文件instance
* console,初化一个console日志instance
*config:{
     "log_path":"日志路径",
     "log_name"："日志名",
     "log_level":"日志等级", //debug,trace,info,warn,error,fatal
     "log_chan_size": 日志通道数量,
     "log_split_type":"日志切隔类型", //hour(按时间切割),size(按容量切割)
 }
 对于console日志实例只需要配置level即可
*/

func InitLogger(name string,config map[string] string)(err error){

	switch name {
	case "file":
		log,err = NewFileLogger(config)
	case "console":
		log,err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s",name)
	}
	return
}


func Debug(format string, args...interface{}){
	log.Debug(format,args...)
}

func Trace(format string, args...interface{}){
	log.Trace(format,args...)
}

func Info(format string, args...interface{}){
	log.Info(format,args...)
}

func Warn(format string, args...interface{}){
	log.Warn(format,args...)
}

func Error(format string, args...interface{}){
	log.Error(format,args...)
}

func Fatal(format string, args...interface{}){
	log.Fatal(format,args...)
}

func Close(){
	log.Close()
}

package logger

//日志接口

type ILogger interface {

	Init() //初始化日志设备
	SetLevel(level int) //设置日志等级
	Debug(format string, args...interface{})
	Trace(format string, args...interface{})
	Info(format string, args...interface{})
	Warn(format string, args...interface{})
	Error(format string, args...interface{})
	Fatal(format string, args...interface{})
	Close()
}

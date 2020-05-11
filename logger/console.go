package logger

import (
	"cache-system/constant"
	"cache-system/cuserror"
	"cache-system/util"
	"fmt"
	"os"
)

func NewConsoleLogger(config map[string] string) (iLog ILogger,err error){
	logLevel,ok := config["log_level"]
	if !ok {
		err = cuserror.ErrNotFoundLevel
		return
	}

	level := util.GetLogLevel(logLevel)
	iLog = &ConsoleLogger{level:level}
	return
}


type ConsoleLogger struct{
	level int
}

func (c *ConsoleLogger) Init() {
	
}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < constant.LogLevelDebug || level >  constant.LogLevelFatal{
		level = constant.LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > constant.LogLevelDebug{
		return
	}
	logData := util.WriteLog(constant.LogLevelDebug,format,args...)
	fmt.Fprintf(os.Stdout,"%s %s (%s:%s:%d) %s\n",logData.TimeStr,logData.LevelStr,logData.Filename,logData.FuncName,logData.LineNo,logData.Message)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > constant.LogLevelTrace{
		return
	}
	logData := util.WriteLog(constant.LogLevelTrace,format,args...)
	fmt.Fprintf(os.Stdout,"%s %s (%s:%s:%d) %s\n",logData.TimeStr,logData.LevelStr,logData.Filename,logData.FuncName,logData.LineNo,logData.Message)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > constant.LogLevelInfo{
		return
	}
	logData := util.WriteLog(constant.LogLevelInfo,format,args...)
	fmt.Fprintf(os.Stdout,"%s %s (%s:%s:%d) %s\n",logData.TimeStr,logData.LevelStr,logData.Filename,logData.FuncName,logData.LineNo,logData.Message)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > constant.LogLevelWarn{
		return
	}
	logData := util.WriteLog(constant.LogLevelWarn,format,args...)
	fmt.Fprintf(os.Stdout,"%s %s (%s:%s:%d) %s\n",logData.TimeStr,logData.LevelStr,logData.Filename,logData.FuncName,logData.LineNo,logData.Message)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > constant.LogLevelError{
		return
	}
	logData := util.WriteLog(constant.LogLevelError,format,args...)
	fmt.Fprintf(os.Stdout,"%s %s (%s:%s:%d) %s\n",logData.TimeStr,logData.LevelStr,logData.Filename,logData.FuncName,logData.LineNo,logData.Message)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > constant.LogLevelFatal{
		return
	}
	logData := util.WriteLog(constant.LogLevelFatal,format,args...)
	fmt.Fprintf(os.Stdout,"%s %s (%s:%s:%d) %s\n",logData.TimeStr,logData.LevelStr,logData.Filename,logData.FuncName,logData.LineNo,logData.Message)

}

func (c *ConsoleLogger) Close() {

}

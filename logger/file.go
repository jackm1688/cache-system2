package logger

import (
	"cache-system/constant"
	"cache-system/cuserror"
	"cache-system/util"
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileLogger struct {
	level         int
	logPath       string
	logName       string
	file          *os.File
	warnFile      *os.File
	LogDataChan   chan *util.LogData
	logSplitType  int   //日志切割类型
	logSplitSize  int64 //日志切割大小
	lastSplitHour int
}

func NewFileLogger(config map[string]string) (iLog ILogger, err error) {

	logPath, ok := config["log_path"]
	if !ok {
		err = cuserror.ErrNotFoundPath
		return
	}

	logLevel, ok := config["log_level"]
	if !ok {
		err = cuserror.ErrNotFoundLevel
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = cuserror.ErrNotFoundName
		return
	}

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
		return
	}

	var logSplitType int = constant.LogSplitTypeHour
	var logSplitSize int64
	logSplitTypeStr, ok := config["log_split_type"]
	if !ok {
		logSplitTypeStr = "hour"
	} else {
		if logSplitTypeStr == "size" {
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "104857600" //100MB
			}

			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600 //解析失败，给一个默认值100MB
			}
			logSplitType = constant.LogSplitTypeSize
		} else {
			logSplitType = constant.LogSplitTypeHour
		}
	}

	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000 //转换失败，设置一个默认值50000
	}

	level := util.GetLogLevel(logLevel)

	iLog = &FileLogger{
		level:         level,
		logPath:       logPath,
		logName:       logName,
		LogDataChan:   make(chan *util.LogData, chanSize),
		logSplitType:  logSplitType,
		logSplitSize:  logSplitSize,
		lastSplitHour: time.Now().Hour(),
	}

	iLog.Init()
	return
}

func (f *FileLogger) Init() {
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed:%v", filename, err))
	}
	f.file = file

	//写错误日志和fatal日志的文件
	/**filename = fmt.Sprintf("%s/%s.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed:%v", filename, err))
	}
	f.warnFile = file
	 */

	go f.writeLogBackground() //异步写日志
}

func (f *FileLogger) writeLogBackground() {
	var file *os.File
	for logData := range f.LogDataChan {
		file =  f.file
		//检查是否需要切割日志文件
		f.checkSplitFile()
		fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", logData.TimeStr, logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
	}
}

//func (f *FileLogger) checkSplitFile(warnFile bool)
func (f *FileLogger) checkSplitFile() {

	if f.logSplitType == constant.LogSplitTypeHour {
		f.splitFileByHour()
		return
	}
	f.splitFileBySize()
}

//func (f *FileLogger) splitFileBySize(warnFile bool) v1
func (f *FileLogger) splitFileBySize() {

	file := f.file

	/** v1
	if warnFile {
		file = f.warnFile
	}
	 */

	startInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := startInfo.Size()
	if fileSize <= f.logSplitSize {
		return
	}

	now := time.Now()
	backupFilename, filename := f.backFileInfo(now, constant.LogSplitTypeSize)
	file.Close()
	os.Rename(filename, backupFilename)

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
		fmt.Printf("reopen file faield:%v\n", filename)
		return
	}

	f.file = file
	/** v1
	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
	 */


}
//func (f *FileLogger) backFileInfo(warnFile bool, now time.Time, splitType int) (backupFilename string, filename string)
func (f *FileLogger) backFileInfo(now time.Time, splitType int) (backupFilename string, filename string) {

	if splitType == constant.LogSplitTypeHour {

		backupFilename = fmt.Sprintf("%s/%s.log.%02d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)

	} else {
		backupFilename = fmt.Sprintf("%s/%s.log.%02d%02d%02d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}

	return
}
//func (f *FileLogger) splitFileByHour(warnFile bool) v.1
func (f *FileLogger) splitFileByHour() {

	now := time.Now()
	hour := now.Hour()
	if hour == f.lastSplitHour {
		return
	}

	f.lastSplitHour = hour

	backupFilename, filename := f.backFileInfo(now, constant.LogSplitTypeHour)

	file := f.file

	/** v1
	if warnFile {
		file = f.warnFile
	}
	 */
	f.Close()
	os.Rename(filename, backupFilename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("reopen file faield:%v\n", err)
		return
	}

	f.file = file

	/** v1
	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
	 */
}

func (f *FileLogger) SetLevel(level int) {
	if level < constant.LogLevelDebug || level > constant.LogLevelFatal {
		level = constant.LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > constant.LogLevelDebug {
		return
	}
	logData := util.WriteLog(constant.LogLevelDebug, format, args...)
	joinChan(f.LogDataChan, logData)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > constant.LogLevelTrace {
		return
	}
	logData := util.WriteLog(constant.LogLevelTrace, format, args...)
	joinChan(f.LogDataChan, logData)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > constant.LogLevelInfo {
		return
	}
	logData := util.WriteLog(constant.LogLevelInfo, format, args...)
	joinChan(f.LogDataChan, logData)
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > constant.LogLevelWarn {
		return
	}
	logData := util.WriteLog(constant.LogLevelWarn, format, args...)
	joinChan(f.LogDataChan, logData)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > constant.LogLevelError {
		return
	}
	logData := util.WriteLog(constant.LogLevelError, format, args...)
	joinChan(f.LogDataChan, logData)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > constant.LogLevelFatal {
		return
	}
	logData := util.WriteLog(constant.LogLevelFatal, format, args...)
	joinChan(f.LogDataChan, logData)
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}

func joinChan(logDataChan chan *util.LogData, logData *util.LogData) {
	select {
	case logDataChan <- logData:
	default:
		//队列满了，则丢弃
	}
}

package logger

import (
	"math/rand"
	"testing"
)

func TestConsoleLogger(t *testing.T) {
	config := make(map[string] string)
	config["log_level"] = "debug"
	InitLogger("console",config)

	Debug("this is a %s","test debug")
	Trace("this is a %s","test trace")
	Info("this is a %s","test info")
	Warn("this is a %s","test warn")
	Error("this is a %s","test error")
	Fatal("this is a %s","test fatal")
}

func TestFileLogger(t *testing.T){

	/**
	onfig:{
	     "log_path":"日志路径",
	     "log_name"："日志名",
	     "log_level":"日志等级", //debug,trace,info,warn,error,fatal
	     "log_chan_size": 日志通道数量,
	     "log_split_type":"日志切隔类型", //hour(按时间切割),size(按容量切割)
	 */
	config := make(map[string]string)
	config["log_path"] = "."
	config["log_name"] = "test-mfz"
	config["log_level"] = "debug"
	config["log_chan_size"] = "10"
	config["log_split_type"] = "size"
	config["log_split_size"] = "104857600"

	err := InitLogger("file",config)
	if err != nil {
		t.Errorf("test InitLogger failed:%v",err)
	}


	var i int = 0
	for{
		i++
		Debug("命令,相信大家都不陌生,常见的情况会使用这个命令做单测试、基准测试和http测试。go test还是有很多flag 可以帮助我们做更多的分析,比如测试覆 %s,rand=%d","test debug",i+rand.Intn(100000))
		Trace("命令,相信大家都不陌生,常见的情况会使用这个命令做单测试、基准测试和http测试。go test还是有很多flag 可以帮助我们做更多的分析,比如测试覆f  %s,rand=%d","test trace",i+rand.Intn(200000))
		Info("命令,相信大家都不陌生,常见的情况会使用这个命令做单测试、基准测试和http测试。go test还是有很多flag 可以帮助我们做更多的分析,比如测试覆d aafdfdf %s,rand=%d","test info",i+rand.Intn(300000))
		Warn("命令,相信大家都不陌生,常见的情况会使用这个命令做单测试、基准测试和http测试。go test还是有很多flag 可以帮助我们做更多的分析,比如测试覆 %s,rand=%d","test warn",i+rand.Intn(400000))
		Error("命令,相信大家都不陌生,常见的情况会使用这个命令做单测试、基准测试和http测试。go test还是有很多flag 可以帮助我们做更多的分析,比如测试覆ff %s,rand=%d","test error",i+rand.Intn(500000))
		Fatal("命令,相信大家都不陌生,常见的情况会使用这个命令做单测试、基准测试和http测试。go test还是有很多flag 可以帮助我们做更多的分析,比如测试覆 %s,rand=%d","test fatal",i+rand.Intn(600000))

		//Close()
	}


}

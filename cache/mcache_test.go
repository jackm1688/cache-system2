package cache

import (
	"cache-system/logger"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func init() {
	config := make(map[string]string,1)
	config["log_level"] = "debug"

	logger.InitLogger("console",config)
	InitMCache()
}

func TestSet(t *testing.T) {
	Set("name","mfz",5) //10秒后过期
	time.Sleep(5*time.Second)
	Set("age","22",25) //10秒后过期
	time.Sleep(5*time.Second)
	Set("work","it-dev",40) //10秒后过期
}

func TestGet(t *testing.T) {
	value,ok := Get("name")
	if ok {
		fmt.Printf("name:%v\n",value)
	}else {
		t.Errorf("Get key:%v is not exists","name")
	}
	select {

	}
}



func MStat(t *testing.T)  {

	type MemStatus struct {
		All  uint32 `json:"all"`
		Used uint32 `json:"used"`
		Free uint32 `json:"free"`
		Self uint64 `json:"self"`
	}

	memStat := new(runtime.MemStats)
	//自身占用
	for i:=0; i < 200000; i++{

		runtime.ReadMemStats(memStat)
		mem := MemStatus{}
		mem.Self = memStat.Alloc

		time.Sleep(1*time.Second)
	}

}


package main

import (
	"cache-system/logger"
	"fmt"
	"math/rand"
	"time"
)
import "cache-system/cache"

func init(){
	rand.Seed(time.Now().UnixNano()) //初始化随机种子
}
func InitLog(){
	config := make(map[string]string,1)
	config["log_level"] = "debug"
	logger.InitLogger("console",config)
}

func InitCache(){
	cache.InitMCache()
}

func main(){
	InitLog() //初始化日志实例
	InitCache() //初化缓存实例

	for i:= 0; i < 500; i++{
		cache.Set("name:"+fmt.Sprintf("%d",i+1),"tom:"+fmt.Sprintf("%d",i+1),rand.Int63n(10000))
		//time.Sleep(200*time.Millisecond)
	}

	/*
	time.Sleep(200*time.Millisecond)
	cache.Flush()
	time.Sleep(200*time.Millisecond)
	fmt.Printf("get cache table size:%d\n",cache.Keys())
	time.Sleep(200*time.Millisecond)
    */

	fmt.Printf("get cache table size:%d\n",cache.Keys())
	//查询name100 key的值
	val,ok :=cache.Get("name:100")
	fmt.Printf("get name:100 value:%v, status:%v\n",val,ok)
	//删除name100 key
	cache.Del("name:100")
	//查询name100 key是否删除成功
	fmt.Printf("get name:100 key:%v\n",cache.Exists("name:100"))
	time.Sleep(2*time.Second) //等待两秒后开清理过期任务
	go cache.ClearExpireNode()

	select { //阻塞等待

	}
}

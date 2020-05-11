package cache

import (
	"cache-system/logger"
	"time"
)

var cache ICache

var currentSize int64

func InitMCache()  {
	//默认初始化一个cache实例，并设置可用存储不能过物理内存可用容量和总容量
	logger.Info("init cache instance,cap:%v","256MB")
	cache = NewCache("256MB",5) //可用存储容量
	//go clearExpireNode() //执行清理任务
}


func SetMaxMemory(size string) bool{
	logger.Info("exec cache.SetMaxMemory,value:%v",size)
	return  cache.SetMaxMemory(size)
}

// 设置⼀一个缓存项，并且在expire时间之后过期
func Set(key string, val interface{}, expire int64){

	var ex time.Duration = time.Duration(expire* int64(1000000000))
	logger.Info("exec cache.Set,key:%v,value:%v,expire:%v",key,val,expire)
	cache.Set(key,val,ex)
}

func Get(key string) (interface{}, bool) {
	v,t := cache.Get(key)
	logger.Info("exec cache.Get,in-value:%v,out-value:%v,status:%v",key,v,t)
	return v,t
}

// 删除⼀个值
func Del(key string) bool {
	t := cache.Del(key)
	logger.Info("exec cache.Del, value:%v",key)
	return  t
}

// 检测⼀个值 是否存在
func Exists(key string) bool {
	t := cache.Exists(key)
	logger.Info("exec cache.Exists,value:%v,status:%v",key,t)
	return t
}

// 情况所有值
func Flush() bool {
	t := cache.Flush()
	logger.Info("exec cache.Flush(),ret-value:%v",t)
	return t
}

// 返回所有的key 多少
func Keys() int64 {
	t :=  cache.Keys()
	logger.Info("exec cache.Keys, ret-value:%v",t)
	return t
}

func ClearExpireNode(){
	cache.ClearExpireNode()
}

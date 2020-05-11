package cache

import (
	"cache-system/logger"
	"cache-system/util"
	"sync"
	"time"
)

type Cache struct {
	SizeStr string //1KB，100KB，1MB，2MB，1GB
	size int64
	cookieMap *linkedList //存在cache
	rwLock sync.RWMutex //读写所
	Interval int // 多长时间执行一次清理任务，操作单位秒
}

func NewCache(size string,interval int) ICache{

	return  &Cache{
		SizeStr:   size,
		size:      util.ParseSize(size),
		cookieMap: NewLinkedList(),
		Interval:interval,
	}
}

func (c *Cache)SetMaxMemory(size string) bool  {
	c.rwLock.RLock()
	defer c.rwLock.Unlock()
	c.SizeStr = size
	c.size = util.ParseSize(size)
	return true
}


// 设置⼀一个缓存项，并且在expire时间之后过期
func (c *Cache)Set(key string, val interface{}, expire time.Duration) {
	c.rwLock.Lock() //使用互斥锁来保证写的数据确证性
	defer c.rwLock.Unlock()
	if c.Exists(key){
		c.cookieMap.Delete(key)
	}

	c.cookieMap.Add(key,val,expire) //Add方法在头部插入
}

func (c *Cache)Get(key string) (interface{}, bool){
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	var t bool = false
	cache := c.cookieMap.Get(key)
	if cache != nil {
		t = true
	}
	if cache != nil {
		return  cache.Value,t
	}
	return nil,false
}

// 删除一个值
func (c *Cache)Del(key string) bool{
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	return  c.cookieMap.Delete(key)
}

// 检测一个值是否存在
func (c *Cache)Exists(key string) bool{
	return  c.cookieMap.IsExists(key)
}

// 清空所有值
func (c *Cache)Flush() bool{
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	return  c.cookieMap.Empty()
}

// 返回所有的key 多少
func (c *Cache)Keys() int64{
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	return  c.cookieMap.Size()
}


func (c *Cache) ClearExpireNode(){
	for{
		logger.Info("exec clear expire key task")
		select {
		case <-time.After(time.Duration(c.Interval) *time.Second):
			if keys :=  c.cookieMap.GetExpireKeys(); len(keys) != 0 {
				for _,key := range keys{
					logger.Info("clear expire key:%v",key)
					c.Del(key)
				}
			}
		}
	}
}



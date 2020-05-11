package cache

import "time"

/**
该程序需要满⾜足以下要求:
1. 支持设定过期时间，精度为秒级。
2. 支持设定最⼤大内存，当内存超出时候做出合理理的处理理。
3. 支持并发安全。
4. 为简化编程细节，无需实现数据落地。
 */

/**
⽀支持过期时间和最⼤大内存⼤大⼩小的的内存缓存库。
*/
type ICache interface {
	//size 是⼀一个字符串串。⽀支持以下参数: 1KB，100KB，1MB，2MB，1GB 等
	SetMaxMemory(size string) bool
	// 设置⼀一个缓存项，并且在expire时间之后过期
	Set(key string, val interface{}, expire time.Duration) // 获取⼀一个值
	Get(key string) (interface{}, bool)
	// 删除⼀一个值
	Del(key string) bool
	// 检测⼀一个值 是否存在
	Exists(key string) bool
	// 情况所有值
	Flush() bool
	// 返回所有的key 多少
	Keys() int64
	//清理过期的节点
	ClearExpireNode()
}

package cache


import (
	"fmt"
	"time"
)

/**
 * 缓存存在结构(使用单链表实现)
 */
type mNode struct{
	Key string  //键
	Value interface{} //值
	InsertTime time.Time //记录插入时间
	Expire time.Duration //过期间隔(单位为秒)
	Next *mNode //定义一地址域(指向下一个)
}

type linkedList struct {
	firstNode *mNode //头结点
}

func NewLinkedList() *linkedList{
	return &linkedList{}
}

//判断是否为空的单链接表
func (link *linkedList) isEmpty() bool {
	if link.firstNode == nil {
		return true
	}
	return false
}


//返回总数
func (link *linkedList) Size() int64{
	first := link.firstNode //获取链表的头结点
	var count  int64 = 0 //定义一个计数器
	for first != nil {
		//如果节点不为空，则count++
		count++
		first = first.Next
	}
	return count
}

//从表头添加元素
func (link *linkedList)Add(key string,value interface{},expire time.Duration) bool{

	node := &mNode{
		Key:key,
		Value: value,
		Expire:expire,
		InsertTime:time.Now(),
	}
	node.Next = link.firstNode
	link.firstNode = node
	return true
}

//在尾部添加数据,需要从头部开始遍历，直到nil
func (link *linkedList) Append(key string,value interface{},expire time.Duration) bool{
	newNode := &mNode{
		Key:key,
		Value:value,
		Expire:expire,
		InsertTime:time.Now(),
	}
	node := link.firstNode
	//首部是空
	if node == nil {
		link.firstNode = newNode
		return  true
	}else {
		for node.Next != nil {
			node = node.Next
		}
		//已经到最后
		node.Next = newNode
		return  true
	}
	return  false
}

//在指定位置插入
func (link *linkedList)Insert(index int64,key string,value interface{},expire time.Duration)  bool {

	newMode := &mNode{
		Key:key,
		Value:value,
		Expire:expire,
	}
	node := link.firstNode
	if index < 0 {
		//index小于0就放在首部
		link.Add(key,value,expire)
		return  true
	} else if index > link.Size(){
		//index大于 长度就放在尾部
		link.Append(key,value,expire)
		return  true
	}else {
		var count int64 = 0
		//找到index之前的元素
		for count < (index -1 ){
			node = node.Next
			count +=1
		}
		//已经找到index之前的元素
		newMode.Next = node.Next
		node.Next = newMode
		return  true
	}
	return  false
}

//删除指定元素，从首部遍历该元素删除，并且需要维护指针
func (link *linkedList) Delete(key interface{})  bool {

	node := link.firstNode
	//如果是首部
	if node != nil && node.Key == key {
		link.firstNode = node.Next
	}else{
		for node != nil  &&  node.Next != nil {
			//找到，改指针
			if node.Next.Key == key{
				node.Next = node.Next.Next
				return  true
			}else{
				node = node.Next
			}
		}
	}
	return false
}

//循环遍历链表
func (link *linkedList) forEachLink()  {

	node := link.firstNode
	for node != nil  {
		str := fmt.Sprintf("{\"%v\":%v}",node.Key,node.Value)
		fmt.Printf("%v\n",str)
		node = node.Next
	}
}


func (link *linkedList) Get(key string) *mNode  {

	node := link.firstNode
	for node != nil  {
		if node.Key == key {
			return  node
		}
		node = node.Next
	}
	return nil
}


func (link *linkedList) IsExists(key string) bool  {

	node := link.firstNode
	for node != nil  {
		if node.Key == key {
			return  true
		}
		node = node.Next
	}
	return false
}

//检查这个key是否过期
func (link *linkedList) isExpire(key string) (string,bool)  {

	node := link.firstNode
	for node != nil  {
		if node.Key == key {
			return key,time.Now().Sub(node.InsertTime) > node.Expire
		}
		node = node.Next
	}
	return key,false
}

func (link *linkedList) Empty() bool{
	node := link.firstNode
	for node != nil  {
		link.Delete(node.Key)
		node = node.Next
		//node = nil
	}

	return true
}

//获取所有过期的keys
func (link *linkedList) GetExpireKeys() (keys []string ) {

	node := link.firstNode
	for node != nil  {
		if time.Now().Sub(node.InsertTime) > node.Expire {
			keys = append(keys,node.Key)
		}
		node = node.Next
	}
	return
}




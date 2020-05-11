package cache_server

import (
	"cache-system/logger"
	"fmt"
	"net"
)

type CacheServer struct {
	Addr string
	Conn map[string]*net.Conn
}

func NewCacheServer() *CacheServer{
	return  &CacheServer{
		Conn: make(map[string]*net.Conn),
	}
}

func (cs *CacheServer)Server()  {

	logger.Info("start cache server:%s",cs.Addr)
	server,err := net.Listen("tcp",cs.Addr)
	if err != nil {
		panic(fmt.Errorf("cache server start lisen failed:%v",err))
	}
	for{
		conn,err := server.Accept()
		if err != nil {
			logger.Error("get client connection failed:%v",err)
			continue
		}
	}
}



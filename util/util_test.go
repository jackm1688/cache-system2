package util

import (
	"fmt"
	"runtime"
	"testing"
)

func TestParseSize(t *testing.T) {

	config := make(map[string]string,1)
	config["log_level"] = "debug"
	ParseSize("100KB")
	a := "我爱你"
	fmt.Printf("a:%d\n",len(a))
	fmt.Println("----------------------",runtime.MemStats{}.HeapAlloc)
}

package main

import (
	"fmt"
	"time"
)

// SafeCounter 的并发使用是安全的。
type NotSafeCounter struct {
	v   map[string]int
}

// Inc 增加给定 key 的计数器的值。
func (c *NotSafeCounter) Inc(key string) {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
}

// Value 返回给定 key 的计数器的当前值。
func (c *NotSafeCounter) Value(key string) int {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	return c.v[key]
}

func main() {
	c := NotSafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

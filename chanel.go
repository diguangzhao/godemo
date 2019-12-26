package main

import (
	"fmt"
	"time"
)

func main() {
	exp3()
}
// exp3 用来测试访问一个已经关闭的且里面还有值未取出的 chanel 会发生什么事？
// 结果是先取出 chanel 里面的值，之后返回零值
func exp3(){
	sigCh:= make(chan string, 2)
	ch:=make (chan int,2)
	ch <- 3
	ch <- 4
	close(ch)

	go func (){
		for i:=0;i<6;i++ {
			a := <-ch
			fmt.Println("b goroutine recieve: " ,a)
			sigCh <-"over"
		}
	}()
	time.Sleep(1 * time.Second)
	for i:=0;i<6;i++  {
		fmt.Println("sigCh", <-sigCh)
		time.Sleep(1 * time.Second)
	}
}
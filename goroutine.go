package main

/*
* @Author: guangzhao.di
* @Date:   2019-12-19 21:11:46
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-19 22:06:43
*/

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan bool)
	go func() {
		fmt.Println("closure func")
		chanel <- true
		time.Sleep(2 * time.Second)
		chanel <- false
		close(chanel)
	}()
	for range chanel  {
		fmt.Println("chanel", chanel)
	}
}
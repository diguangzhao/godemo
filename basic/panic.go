package main
/*
* @Author: guangzhao.di
* @Date:   2019-12-20 00:17:18
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-20 00:21:09
*/

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("run funcA")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B", err)
		}
	}()
	panic("panic in B")
}

func C() {
	fmt.Println("run funcC")
}
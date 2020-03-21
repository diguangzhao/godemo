package main

/*
* @Author: guangzhao.di
* @Date:   2019-12-19 21:11:46
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-19 22:06:43
*/

import (
	"fmt"
	"sync"

	// "sort"
)

type person struct {
	Name string
	Age int
}
func main() {
	sync.Mutex{}.Lock()
	a := person{"张三", 2}
	b := person{}
	b = a
	fmt.Println(a.Name, a.Age, &a)
	fmt.Printf("a %p \n", &a)
	fmt.Printf("b %p \n", &b)
}
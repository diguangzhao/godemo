package main

/*
* @Author: guangzhao.di
* @Date:   2019-12-19 21:11:46
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-19 22:06:43
*/

import (
	"fmt"
	// "sort"
)

type structA struct {
	A string
}

type interfaceA interface {
	 Exec(a *structA) error
}

type structB struct {
	B string
}

func (this 	structB) Exec(a *structA) error  {
	fmt.Println("structB", this)
	this.B = "abc"
	return nil
}

func main() {
	var inter interfaceA = new(structB)
	fmt.Println("inter", inter)
	inter.Exec(&structA{})
	fmt.Println("inter", inter)

	var inter2 interfaceA = &structB{}
	fmt.Println("inter2", inter2)
	inter2.Exec(&structA{})
	fmt.Println("inter2", inter2)

	var inter3 interfaceA = structB{}
	fmt.Println("inter3", inter3)
	inter3.Exec(&structA{})
	fmt.Println("inter3", inter3)
}
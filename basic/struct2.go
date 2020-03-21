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

type A struct {
	 Aname string
	 Aint int
	 Aint2 int
}

type Acopy struct {
	Aname string
	Aint int
	Aint2 int
}

type B struct {
	Bname string
	Bint int
	Bint2 int
}

type AB struct {
	A
	B
	Aint int
}

type ABC struct {
	A
	B
	ABCname string
	ABCint int
}

func main() {
	a := &A{}
	ab := Acopy{}

	ab = Acopy(*a)

	fmt.Println(ab)
	test(a)
}

func test(a *A)  {
	fmt.Println(a)
}
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

type M struct {
	Name string
}

func main() {
	m := M{}
	m.Test(1)
	(*M).Test(&M{}, 1)
}

func (this *M) Test(a int)  {
	fmt.Println(a)
}
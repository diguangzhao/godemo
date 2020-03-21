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

func main() {
	// a := []byte{'a','b', 'c', 'd', 'e', 'f', 'g'}
	// sa := a[3:6]
	// fmt.Println(string(sa))

	// sb := sa[2:3]
	// fmt.Println(sb)

	// s1 := [...]int{1,2,3,4,5}
	// s2 := s1[]
	// s2[1] = 5;
	// fmt.Println(s1)

	// s1 := []int{6,1,2,3,4,5}
	// sort.Ints(s1)
	// fmt.Println(s1)

	s1 := [...]int{1,2,3,4,5}
	s2 := s1[2:3]
	copy(s1, s2)
	fmt.Println(s1)
}
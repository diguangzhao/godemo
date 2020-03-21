package main

import "fmt"
/*
* @Author: guangzhao.di
* @Date:   2019-12-19 23:52:19
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-19 23:59:32
*/

func main() {
	a := 1;
	fmt.Printf("main 	%p \n", &a)
	fa := closure(a)
	fa(2)

	A(a)
}

func closure(a int) func(int)int{
	fmt.Printf("closure 	%p \n", &a)
	return func(b int) int {
		fmt.Printf("closure return	%p \n", &a)
		return a + b
	}
}

func A(a int) {
	fmt.Printf("A 	%p \n", &a)
}
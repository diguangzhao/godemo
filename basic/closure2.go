package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println("i", i)
		}()
		defer A(i)
	}
}

func A(a int) {
	fmt.Println("a", a)
}
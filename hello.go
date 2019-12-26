package main

import "fmt"

/*
6 0110
*/
func main() {
	a := 1
	LABEL:
	for a=0;  a < 3; a++ {
		fmt.Println(a)
		for {
			goto LABEL
		}
	}
	fmt.Println(a)

}

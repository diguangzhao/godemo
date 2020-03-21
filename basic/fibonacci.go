package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var n,s1,s2 = 0,1,1
	return func() int {
		switch n {
		case 0:
			n++
			return 0
		case 1,2:
			n++
			return 1
		default:
			s1, s2 = s2, s1+s2
			return s2
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	f1 := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f1())
	}
}

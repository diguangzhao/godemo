package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	foo := "bar"
	if true {
		foo := "baz"
		doSomething(foo)
	}
	fmt.Println(foo)
}
func doSomething(foo string)  {
	json.Unmarshal()
}
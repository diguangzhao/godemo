package main

import (
	"fmt"
	"sync"
)

func main()  {
	//test()
	test2()
}

func test()  {
	w := sync.WaitGroup{}

	for i:=0; i<10; i++ {
		w.Add(1)
		go func(i int) {
			w.Done()
			fmt.Println("i", i)
		}(i)
	}
	w.Wait()
}

func test2()  {
	w := sync.WaitGroup{}

	for i:=0; i<10; i++ {
		w.Add(1)
		go getPrint(i, &w)()
	}
	w.Wait()
}

func getPrint(i int, group *sync.WaitGroup) func()  {

	return func() {
		group.Done()
		fmt.Println("i##", i)
	}
}
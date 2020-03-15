package main

import "fmt"

var heapsLen int
var heaps []int
var count   = 0

func BuildHeaps(arr []int)  {
	heapsLen = len(arr)
	for i:=0; i<heapsLen; i++ {
		heapify(arr[i], i)
	}
	fmt.Println("origin heap:", heaps)
}

func heapify(num int, i int)  {
	heaps = append(heaps, num)
	adjustHeaps(i)
}

func adjustHeaps(i int)  {
	if i>1 && heaps[i] > heaps[(i-1)/2] {
		Swap(&heaps[i], &heaps[(i-1)/2])
		//fmt.Println("heaps:", heaps)
		adjustHeaps((i-1)/2)
	}
	return
}

func Swap(a *int, b *int)  {
	*a, *b = *b, *a
	count++
}

func adjustTop(i int)  {
	if i >= heapsLen {
		return
	}
	fmt.Println("heaps:", heaps)
	if (2 * i +1 < heapsLen && heaps[i] < heaps[2 * i +1]) || (2 * i +2 <heapsLen && heaps[i] < heaps[2 * i + 2]) {
		if 2 * i +2 <heapsLen && heaps[2 * i +1] < heaps[2 * i +2] {
			Swap(&heaps[i], &heaps[2 * i +2])
			adjustTop(2 * i +2)

		} else {
			Swap(&heaps[i], &heaps[2 * i +1])
			adjustTop(2 * i +1)
		}
	}
	return
}


func pop() int {

	fmt.Println("pre heaps:", heaps)

	Swap(&heaps[0], &heaps[heapsLen-1])
	fmt.Println("swap heaps:", heaps, "heapsLen", heapsLen)

	var num int =heaps[heapsLen-1]
	heaps = heaps[0:heapsLen-1]
	heapsLen--
	fmt.Println("cut heaps:", heaps)
	adjustTop(0)
	fmt.Println("pop:", num, heaps)

	return num
}

func HeapSort(arr []int) []int  {
	BuildHeaps(arr)
	arr = []int{}
	for heapsLen>0 {
		arr = append(arr, pop())
	}
	return arr
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7};
	arr = HeapSort(arr)
	fmt.Println("arr:", arr)
	fmt.Println("count:", count, "len", len(arr))
}

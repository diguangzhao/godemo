package main

import "fmt"

/*
时间复杂度 O(nlog2n)
空间复杂度 O(1)
不稳定的
*/
func QuickSort(slice []int) {
	fmt.Printf("slice%p \n", slice)
	length := len(slice)
	if length <= 1 {
		return
	}
	current := slice[0]
	left := []int{}
	right := []int{}

	for i := 1; i < length; i++ {
		if slice[i] <= current {
			left = append(left, slice[i])
		} else {
			right = append(right, slice[i])
		}
	}
	if len(left) > 1 {
		QuickSort(left)
	}
	if len(right) > 1 {
		QuickSort(right)
	}

	//这里直接赋给slice不行，会重新分配指针，必须copy why？
	tmp := append(append(left, current), right...)
	copy(slice, tmp)
	fmt.Printf("slice%p \n", slice)
}

func main() {

	slice := []int{3, 2, 4, 5, 8}
	fmt.Println("old slice", slice)
	QuickSort(slice)
	fmt.Println("sorted slice", slice)
}

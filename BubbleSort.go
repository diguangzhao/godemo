package main

import "fmt"

func Swap(slice []int, i, j int)  {
	tmp := slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
}
/*
时间复杂度 O(n^2)
空间复杂度 O(1)
稳定的
 */
func BubbleSort(slice []int)  {
	length := cap(slice)
	fmt.Println(length)
	for i:=0; i<length; i++ {
		for j:=length-1; j>i; j-- {
			if (slice[j] < slice[j-1]) {
				Swap(slice, j, j-1)
			}
		}
	}
}

func main() {
	slice := []int{3, 2, 4, 5, 8}
	fmt.Println("old slice", slice)
	BubbleSort(slice)
	fmt.Println("sorted slice", slice)
}

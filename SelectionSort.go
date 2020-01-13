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
不稳定
 */
func SelectionSort(slice []int)  {
	length := cap(slice)

	for i:=0; i<length; i++ {
		minIndex:=i
		for j:=i+1; j<length;j++ {
			if (slice[j] < slice[minIndex]) {
				minIndex = j
			}
		}
		Swap(slice, i, minIndex)
	}
}

func main() {
	slice := []int{3, 2, 4, 5, 8}
	fmt.Println("old slice", slice)
	SelectionSort(slice)
	fmt.Println("sorted slice", slice)
}

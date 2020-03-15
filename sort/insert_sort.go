package main

import "fmt"

func insertSort(arr []int) {
	len := len(arr)
	for i:=1; i<len; i++ {
		if (arr[i] < arr[i-1]) {
			temp := arr[i]
			j:=i
			for ; j>0;j--  {
				if arr[j-1] > temp {
					arr[j] = arr[j-1]
				} else {
					break
				}
			}
			arr[j] = temp

		}
	}
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7,3};
	insertSort(arr)
	fmt.Println("arr:", arr)
}

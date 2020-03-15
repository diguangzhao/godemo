package main

import "fmt"


/*
时间复杂度:
平均时间复杂度:O(n log n) 最好时间复杂度:O(n log n) 最坏时间复杂度:O(n2) 空间复杂度:
O(log n)
稳定性: 不稳定排序
*/
func QuickSort( arr []int ) ([]int) {
	length := len(arr)
	leftArr := []int{}
	rightArr := []int{}
	flag := arr[0]
	for i:=1; i<length; i++ {
		if arr[i] <= flag {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	if len(leftArr)>0 {
		leftArr = QuickSort(leftArr)
	}
	if len(rightArr)>0 {
		rightArr = QuickSort(rightArr)
	}
	arr = append(append(leftArr, flag), rightArr...)
	return arr
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7,3, 4};
	arr = QuickSort(arr)
	fmt.Println("arr:", arr)
	fmt.Println("len:", len(arr))
}

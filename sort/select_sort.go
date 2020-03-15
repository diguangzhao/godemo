package main

import "fmt"

var count int

func selectSort(arr []int) (int) {
	len := len(arr)
	for i:=0; i<len-1; i++ {
		minIndex:=i
		for j:=i+1; j<len;j++  {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		Swap(&arr[minIndex], &arr[i])
	}
	return count
}

func Swap(a *int, b *int)  {
	*a, *b = *b, *a
	count++
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7,3};
	count:=selectSort(arr)
	fmt.Println("arr:", arr)
	fmt.Println("count:", count, "len:", len(arr))
}

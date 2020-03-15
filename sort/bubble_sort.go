package main

import "fmt"

var count  int

func BubbleSort(arr []int) (int) {
	len := len(arr)
	for i:=0; i<len-1; i++ {
		for j:=len-1; j>i;j--  {
			if arr[j] < arr[j-1] {
				Swap(&arr[j], &arr[j-1])
			}
		}
	}
	return count
}

func Swap(a *int, b *int)  {
	*a, *b = *b, *a
	count++
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7,3};
	count:=BubbleSort(arr)
	fmt.Println("arr:", arr)
	fmt.Println("count:", count, "len:", len(arr))
}

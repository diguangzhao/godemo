package main

import "fmt"

func ShellSort(arr []int)  {
	len := len(arr)
	gap := len/2

	for ;gap>=1 ;gap /= 2  {
		for i:=gap; i<len; i++ {
			temp := arr[i]
			j:=i-gap
			for ; j>=0 && arr[j] > temp; j-=gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = temp
		}
	}
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7,3};
	ShellSort(arr)
	fmt.Println("arr:", arr)
}

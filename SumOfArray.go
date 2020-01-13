package main

import "fmt"

func SumOfArray(arr []int, sum int)  {
	var (
		i = 0
		j = len(arr)-1
		last = j
	)
	for i <= j  {
		if arr[i] + arr[j] == sum {
			fmt.Printf("符合：%n,%n \n", i, j)
			if i+1 < last && arr[i+1] == arr[i] {
				i++
				fmt.Printf("符合：%n,%n \n", i, j)
			}
			if j-1>0 && arr[j-1] == arr[j] {
				j--
				continue
			}
			if (i+1 < last && arr[i+1] != arr[i])  && (j-1>0 && arr[j-1] != arr[j]) {
				i ++
				continue
			}
		}
		if arr[i] + arr[j] > sum {
			j --
		}
		if arr[i] + arr[j] < sum {
			i ++
		}
	}
}

func main() {
	slice := []int{1, 3, 3, 4, 5, 8, 9, 10, 11, 12, 18, 20, 20, 20, 21}
	sum := 21
	SumOfArray(slice, sum)
}

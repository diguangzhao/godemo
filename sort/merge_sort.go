package main

import "fmt"

func MergeSort(arr []int)  {
	sort(arr, 0, len(arr) -1)
}

func sort(arr []int, start int, end int)  {

	if start == end {
		return
	}
	mid := (start + end) /2
	sort(arr, start, mid)
	sort(arr, mid + 1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start int, mid int, end int)  {
	temp := []int{}
	lp :=start
	rp :=mid+1
	for lp<=mid || rp<=end {
		if lp >mid {
			temp = append(temp, arr[rp])
			rp++
			continue
		}
		if rp >end {
			temp = append(temp, arr[lp])
			lp++
			continue
		}
		if lp > mid && rp>end {
			break
		}
		if arr[lp] <= arr[rp] {
			temp = append(temp, arr[lp])
			lp++
		} else {
			temp = append(temp, arr[rp])
			rp++
		}
	}

	for i:=0; i<len(temp); i++ {
		arr[start+i] = temp[i]
	}
}

func main() {
	arr := []int{9,1,5,6,4,10,5,8,7,3};
	MergeSort(arr)
	fmt.Println("arr:", arr)
}

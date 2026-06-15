package main

import "fmt"

func bubbleSort(arr []int) []int {

	for pass := 0; pass < len(arr)-1; pass++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{5, 1, 4, 3}))
}

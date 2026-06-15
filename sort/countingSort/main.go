package main

import "fmt"

func countingSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	count := make([]int, max+1)

	for _, n := range nums {
		count[n]++
	}

	result := make([]int, 0, len(nums))

	for num, freq := range count {
		for freq > 0 {
			result = append(result, num)
			freq--
		}
	}

	return result
}

func main() {
	// max = 8
	// count 8 => {0 ,1 ,2 ,2 ,1 ,3 ,1 ,0 ,1}
	//             0  1  2  3  4  5  6  7  8
	fmt.Println(countingSort([]int{3, 2, 1, 2, 3, 4, 5, 6, 8, 5, 5}))
}

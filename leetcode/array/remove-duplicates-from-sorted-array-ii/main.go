package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	index := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}

	return index

}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}

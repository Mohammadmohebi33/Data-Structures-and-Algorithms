package main

import "fmt"

func leftRightDifference(nums []int) []int {

	var res []int

	for i := 0; i < len(nums); i++ {

		sumleft := 0
		sumright := 0

		for j := i + 1; j < len(nums); j++ {
			sumright += nums[j]
		}

		for j := i - 1; j >= 0; j-- {
			sumleft += nums[j]
		}

		diff := sumleft - sumright
		if diff < 0 {
			diff = -diff
		}
		res = append(res, sumleft-sumright)
	}
	return res
}

func main() {
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3}))
}

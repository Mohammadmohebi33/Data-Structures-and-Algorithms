package main

import "fmt"

func runningSum(nums []int) []int {
	var res []int
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res = append(res, sum)
	}
	return res
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4})) // [1,3,6,10]

}

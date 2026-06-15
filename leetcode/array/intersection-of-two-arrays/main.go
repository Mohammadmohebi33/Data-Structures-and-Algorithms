package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	for i := 0; i < len(nums1); i++ {
		if contains(nums2, nums1[i]) && !contains(res, nums1[i]) {
			res = append(res, nums1[i])
		}
	}
	return res
}

func contains(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

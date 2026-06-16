package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	pointer1 := m - 1
	pointer2 := n - 1
	current := m + n - 1

	for pointer2 >= 0 {
		if pointer1 >= 0 && nums1[pointer1] > nums2[pointer2] {
			nums1[current] = nums1[pointer1]
			pointer1--
		} else {
			nums1[current] = nums2[pointer2]
			pointer2--
		}
		current--
	}

}

func main() {

}

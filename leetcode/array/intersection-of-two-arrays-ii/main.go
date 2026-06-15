package main

func intersect(nums1 []int, nums2 []int) []int {
	freq := make(map[int]int)
	res := []int{}

	for _, num := range nums1 {
		freq[num]++
	}

	for _, num := range nums2 {
		if freq[num] > 0 {
			res = append(res, num)
			freq[num]--
		}
	}

	return res
}
func main() {

}

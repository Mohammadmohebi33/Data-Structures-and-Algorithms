package main

func moveZeroes(nums []int) {
	insertPos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[insertPos] = nums[i]
			insertPos++
		}
	}

	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}
}

func sortColors(nums []int) {
	for pass := 0; pass < len(nums)-1; pass++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

func main() {

	moveZeroes([]int{0, 1, 0, 3, 12})

}

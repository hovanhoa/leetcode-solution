
func moveZeroes(nums []int) {
	l := 0
	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l += 1
		}
	}

	for l < len(nums) {
		nums[l] = 0
		l += 1
	}
}

func increasingTriplet(nums []int) bool {
	one, two := math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v <= one {
			one = v
		} else if v <= two {
			two = v
		} else {
			return true
		}
	}

	return false
}

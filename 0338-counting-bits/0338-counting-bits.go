func countBits(n int) []int {
	ones := make([]int, n+1)
	for x := 0; x <= n; x++ {
        ones[x] = ones[x>>1] + x&1
	}
	return ones
}
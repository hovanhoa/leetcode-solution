func hammingWeight(n int) int {
    ans := 0
    for n > 0 {
        ans += n % 2
        n /= 2
    }

    return ans
}
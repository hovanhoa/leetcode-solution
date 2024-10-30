func isHappy(n int) bool {
    nums := map[int]bool{}
    for !nums[n] {
        nums[n] = true
        n = sumOfSquare(n)
        
        if n == 1 {
            return true
        }
    }

    return false
}

func sumOfSquare(n int) int {
    ans := 0
    for n > 0 {
        tmp := n % 10
        ans += tmp * tmp
        n /= 10
    }

    return ans
}
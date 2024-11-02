func isUgly(n int) bool {
    if n < 1 {
        return false
    }

    for _, v := range []int{2, 3, 5} {
        for n % v == 0 {
            n /= v
        }
    }

    return n == 1 
}
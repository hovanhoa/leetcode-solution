func climbStairs(n int) int {
    one, two := 1, 0 
    for i := 0; i < n; i++ {
        temp := one
        one = one + two
        two = temp
    }

    return one
}
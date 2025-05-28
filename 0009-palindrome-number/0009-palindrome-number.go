func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    n := 0 
    temp := x
    for temp > 0 {
        n = n*10 + temp%10
        temp /= 10
    }

    return x == n
}
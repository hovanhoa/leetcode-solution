func isHappy(n int) bool {
    m := map[int]struct{}{}
    for {
        m[n] = struct{}{}
        n = sumOfNum(n)

        if n == 1 {
            return true
        }

        if _, ok := m[n]; ok {
            return false
        }
    }

    return false
}

func sumOfNum(n int) int {
    res := 0
    for n > 0 {
        mod := n%10 
        res += mod * mod
        n /= 10
    }

    return res
}
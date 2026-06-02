func isHappy(n int) bool {
    m := map[int]bool{}
    for {
        if _, ok := m[n]; ok {
            return false
        }

        s := 0
        m[n] = true
        for n != 0 {
            s += (n%10)*(n%10)
            n /= 10
            fmt.Println(s, n)
        }

        if s == 1 {
            return true
        }

        n = s
    }

    return false
}
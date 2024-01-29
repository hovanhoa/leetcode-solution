func tribonacci(n int) int {
    if n == 0 {
        return 0
    } else if n < 3 {
        return 1
    }

    t0 := 0
    t1 := 1
    t2 := 1
    
    for i := 0; i < (n - 2); i++ {
        newT := t0 + t1 + t2
        t0 = t1
        t1 = t2
        t2 = newT
    }

    return t2
}
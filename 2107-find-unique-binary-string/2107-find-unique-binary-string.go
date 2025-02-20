func findDifferentBinaryString(nums []string) string {
    m := map[string]bool{}
    for _, num := range nums {
        m[num] = true
    }

    n := len(nums[0])
    for i := 0; i < int(math.Pow(2, float64(n))); i++ {
        s := binToString(i, n)
        if !m[s] {
            return s
        }
    }

    return ""
}

func binToString(n int, l int) string {
    res := ""
    for l > 0 {
        if n%2 == 0 {
            res = "0" + res
        } else {
            res = "1" + res
        }

        n /= 2
        l -= 1
    }

    return res
}
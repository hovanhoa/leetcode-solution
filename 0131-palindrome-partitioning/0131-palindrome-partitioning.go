func partition(s string) [][]string {
    ans, part := [][]string{}, []string{}

    var dfs func(i int)
    dfs = func(i int) {
        if i >= len(s) {
            tmp := make([]string, len(part))
            copy(tmp, part)
            ans = append(ans, tmp)
            return
        }

        for j := i; j < len(s); j++ {
            if isPalindrome(s, i, j) {
                part = append(part, s[i:j+1])
                dfs(j+1)
                part = part[:len(part)-1]
            }
        }
    }

    dfs(0)
    
    return ans
}

func isPalindrome(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }

        l, r = l + 1, r - 1
    }

    return true
}

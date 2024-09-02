func isSubsequence(s string, t string) bool {
    idx := 0
    for i := 0; i < len(t); i ++ {
        if idx < len(s) && s[idx] == t[i] {
            idx ++
        }
    }

    return idx == len(s)
}
func isSubsequence(s string, t string) bool {
    sIndex := 0
    for i := 0; i < len(t); i++ {
        if sIndex == len(s) {
            return true
        }

        if t[i] == s[sIndex] {
            sIndex ++
        }
    }

    return sIndex == len(s)
}
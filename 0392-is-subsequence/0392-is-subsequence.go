func isSubsequence(s string, t string) bool {
    i, j := 0, 0
    for j < len(t) && i < len(s) {
        if s[i] == t[j] {
            i += 1
        }

        j += 1
    }

    return i == len(s)
}
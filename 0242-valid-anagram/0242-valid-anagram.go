func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS, countT := map[byte]int{}, map[byte]int{}
    for i := 0; i < len(s); i++ {
        countS[s[i]] += 1
        countT[t[i]] += 1
    }

    for k, v := range countS {
        if v != countT[k] {
            return false
        }
    }

    return true
}
func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}
    for _, s := range strs {
        k := [26]int{}
        for i := 0; i < len(s); i++ {
            k[s[i]-'a'] += 1
        }

        m[k] = append(m[k], s)
    }

    ans := [][]string{}
    for _, v := range m {
        ans = append(ans, v)
    }

    return ans
}
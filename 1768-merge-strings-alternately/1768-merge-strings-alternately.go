func mergeAlternately(word1 string, word2 string) string {
    ans := []byte{}
    i, j := 0, 0
    for i < len(word1) || j < len(word2) {
        if i < len(word1) {
            ans = append(ans, word1[i])
            i += 1
        }

        if j < len(word2) {
            ans = append(ans, word2[j])
            j += 1
        }
    }    

    return string(ans)
}

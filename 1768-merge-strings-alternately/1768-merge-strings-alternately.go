func mergeAlternately(word1 string, word2 string) string {
    ans := make([]byte, len(word1) + len(word2))
    count := 0
    i, j := 0, 0
    for i < len(word1) && j < len(word2) {
        if i < len(word1) {
            ans[count] = word1[i]
            i++
            count++
        }

        if j < len(word2) {
            ans[count] = word2[j]
            j++
            count++
        }
    }

    for i < len(word1) {
        ans[count] = word1[i]
        i++
        count++
    }

    for j < len(word2) {
        ans[count] = word2[j]
        j++
        count++
    }

    return string(ans)
}
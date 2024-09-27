func mergeAlternately(word1 string, word2 string) string {
    i, j := 0, 0
    result := []byte{}
    for i < len(word1) || j < len(word2) {
        if i < len(word1) {
            result = append(result, word1[i])
            i += 1
        }

        if j < len(word2) {
            result = append(result, word2[j])
            j += 1
        }
    }

    return string(result)
}
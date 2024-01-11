func mergeAlternately(word1 string, word2 string) string {
    len1 := len(word1)
    len2 := len(word2)
    var result string
    maxLen := len1 
    if len2 > len1 {
        maxLen = len2
    }

    for i := 0; i < maxLen; i ++ {
        if i < len1 {
            result += string(word1[i])
        }

        if i < len2 {
            result += string(word2[i])
        }
    }

    return result
}
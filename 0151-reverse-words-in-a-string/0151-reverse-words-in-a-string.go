func reverseWords(s string) string {
    words := strings.Fields(s)
    l, r := 0, len(words) - 1
    for l < r {
        words[l], words[r] = words[r], words[l]
        l++
        r--
    }

    return strings.Join(words, " ")
}
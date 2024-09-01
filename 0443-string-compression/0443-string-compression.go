func compress(chars []byte) int {
    l, r := 0, 0
    for r < len(chars) {
        j := r + 1
        for ;j < len(chars) && chars[j] == chars[j-1]; j++ {}

        chars[l] = chars[r]
        l++
        if j != r + 1 {
            for _, v := range(strconv.Itoa(j-r)) {
                chars[l] = byte(v)
                l++
            }
        }

        r = j
    }

    chars = chars[:l]
    return len(chars)
}
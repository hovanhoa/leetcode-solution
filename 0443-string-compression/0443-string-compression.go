func compress(chars []byte) int {
    j := 0
    count := 0
    char := chars[0]
    for i := 0; i < len(chars); i++ {
        if char == chars[i] {
            count++
        } else {
            chars[j] = char
            j++

            if count > 1 {
                bCount := []byte(strconv.Itoa(count))
                for _, v := range bCount {
                    chars[j] = v
                    j++
                }
            }

            count = 1
            char = chars[i]
        }
    }

    chars[j] = char
    j++

    if count > 1 {
        bCount := []byte(strconv.Itoa(count))
        for _, v := range bCount {
            chars[j] = v
            j++
        }
    }

    return j
}
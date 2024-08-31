func reverseVowels(s string) string {
    vowelsMap := map[byte]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        'A': true,
        'E': true,
        'I': true,
        'O': true,
        'U': true,
    }

    result := []byte(s)
    left, right := 0, len(s) - 1
    for left < right {
        for left < right && !vowelsMap[result[left]]{
            left += 1
        }

        for left < right && !vowelsMap[result[right]] {
            right -= 1
        }

        result[left], result[right] = result[right], result[left]
        left += 1
        right -= 1
    }   

    return string(result)
}
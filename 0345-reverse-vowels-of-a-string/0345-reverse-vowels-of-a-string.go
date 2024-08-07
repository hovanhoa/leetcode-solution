func reverseVowels(s string) string {
    out := []rune(s)
    vowels := map[rune]bool{
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
    left, right := 0, len(s) - 1

    for left < right {
        for left < right && !vowels[out[left]] {
            left += 1
        }

        for left < right && !vowels[out[right]] {
            right -= 1
        }

        if left < right {
            out[left], out[right] = out[right], out[left]
            left += 1
            right -= 1
        }
    } 

    return string(out)
}
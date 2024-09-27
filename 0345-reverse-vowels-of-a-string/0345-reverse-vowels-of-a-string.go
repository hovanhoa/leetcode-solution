func reverseVowels(s string) string {
    vowels := map[byte]bool{
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

    ans := []byte(s)
    l, r := 0, len(ans) - 1
    for l < r {
        for !vowels[ans[l]] && l < r {
            l++
        }

        for !vowels[ans[r]] && l < r {
            r--
        }

        if l < r {
            ans[l], ans[r] = ans[r], ans[l]
        }

        l++
        r--
    }

    return string(ans)
}
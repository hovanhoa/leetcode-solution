func maxVowels(s string, k int) int {
    vowels := map[byte]bool {
        'a': true,
        'e': true, 
        'i': true,
        'o': true,
        'u': true,
    }

    ans := 0
    sByte := []byte(s)
    for i := 0; i < k; i++ {
        if vowels[sByte[i]] {
            ans++
        }
    }

    cur := ans
    for i := k; i < len(sByte); i++ {
        if vowels[sByte[i]] {
            cur++
        }

        if vowels[sByte[i-k]] {
            cur--
        }

        if cur > ans {
            ans = cur
        }
    }

    return ans
}
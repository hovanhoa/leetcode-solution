func maxVowels(s string, k int) int {
    res, num := 0, 0
    vowelMap := map[byte]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }

    for i := 0; i < len(s); i++ {
        if i >= k && vowelMap[s[i-k]] {
            num--
        }

        if vowelMap[s[i]] {
            num++
        }

        if num > res {
            res = num
        }
    }

    return res
}
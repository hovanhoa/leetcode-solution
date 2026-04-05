func lengthOfLongestSubstring(s string) int {
    res := 0
    arr := []byte{}
    for i := 0; i < len(s); i++ {
        for j := range arr {
            if arr[j] == s[i] {
                arr = arr[j+1:]
                break
            }
        }

        arr = append(arr, s[i])
        res = max(res, len(arr))
    }

    return res
}
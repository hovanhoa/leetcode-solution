func lengthOfLongestSubstring(s string) int {
    res := 0
    arr := []byte{}
    for _, v := range []byte(s) {
        for j := range arr {
            if arr[j] == v {
                arr = arr[j+1:]
                break
            }
        }

        arr = append(arr, v)
        res = max(res, len(arr))
    }

    return res
}
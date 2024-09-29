func maxOperations(nums []int, k int) int {
    nMap := make(map[int]int)
    for _,v := range nums {
        nMap[v]++
    }

    ans := 0
    for key := range nMap {
        if nMap[k-key] > 0 && k-key != key {
            ans += min(nMap[k-key], nMap[key])
            nMap[key] = 0
            nMap[k-key] = 0
        } else if nMap[k-key] > 0 {
            ans += nMap[key]/2
            nMap[key] = 0
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
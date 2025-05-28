func maxOperations(nums []int, k int) int {
    ans := 0
    m := map[int]int{}
    for _, v := range nums {
        if m[k-v] > 0 {
            ans += 1
            m[k-v] -= 1
        } else {
            m[v] += 1
        }
    }

    return ans
}
func maxOperations(nums []int, k int) int {
    kMap := map[int]int{}
    count := 0
    for _, v := range nums {
        if kMap[k-v] > 0 {
            count++
            kMap[k-v]--
        } else {
            kMap[v]++
        }
    }

    return count
}
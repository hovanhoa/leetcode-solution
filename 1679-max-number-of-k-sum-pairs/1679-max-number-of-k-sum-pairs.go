func maxOperations(nums []int, k int) int {
    kMap := map[int]int{}
    for _, v := range nums {
        kMap[v]++
    }

    count := 0
    for _, v := range nums {
        if kMap[k-v] > 0 {
            kMap[k-v]--
            count++
        }
    }

    return count/2
}
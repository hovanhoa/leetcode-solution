func majorityElement(nums []int) int {
    m := map[int]int{}
    for _, v := range nums {
        m[v] += 1
    }

    max := 0
    res := 0
    for k, v := range m {
        if v > max {
            max = v
            res = k
        }
    }

    return res
}
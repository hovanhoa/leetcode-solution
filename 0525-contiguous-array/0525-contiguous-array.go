func findMaxLength(nums []int) int {
    zero, one := 0, 0
    m := map[int]int{}
    res := 0

    for i, v := range nums {
        if v == 0 {
            zero += 1
        } else {
            one += 1
        }

        diff := zero - one
        if _, ok := m[diff]; !ok {
            m[diff] = i
        }

        if diff == 0 {
            res = zero + one
        } else {
            res = max(res, i - m[diff])
        }
    }

    return res
}
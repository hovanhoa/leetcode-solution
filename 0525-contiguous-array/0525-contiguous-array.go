func findMaxLength(nums []int) int {
    m := map[int]int{}
    zero, one := 0, 0
    res := 0
    for i, v := range nums {
        if v == 0 {
            zero += 1
        } else {
            one += 1
        }

        if zero == one {
            res = zero + one
        }

        diff := zero - one
        if _, ok := m[diff]; ok {
            res = max(res, i - m[diff])
        } else {
            m[diff] = i
        }
    }

    return res
}
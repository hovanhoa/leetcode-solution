func findMaxLength(nums []int) int {
    zeros, ones := 0, 0
    m := map[int]int{}
    res := 0

    for i, v := range nums {
        if v == 0 {
            zeros += 1
        } else {
            ones += 1
        }

        diff := zeros - ones
        if _, ok := m[diff]; !ok {
            m[diff] = i
        }

        if zeros == ones {
            res = zeros + ones
        } else {
            idx := m[diff]
            res = max(res, i - idx)
        }
    }

    return res
}
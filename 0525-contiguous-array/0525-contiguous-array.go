func findMaxLength(nums []int) int {
    zeros, ones := 0, 0
    m := map[int]int{}
    ans := 0

    for i, v := range nums {
        if v == 0 {
            zeros += 1
        } else {
            ones += 1
        }

        if zeros == ones {
            ans = zeros + ones
        }

        diff := zeros - ones
        idx, ok := m[diff]
        if ok {
            ans = max(ans, i - idx)
        } else {
            m[diff] = i
        }
    }

    return ans
}

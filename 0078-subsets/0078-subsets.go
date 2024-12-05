func subsets(nums []int) [][]int {
    var ans [][]int
    var subset []int

    var dfs func(i int)
    dfs = func(i int) {
        if i >= len(nums) {
            cpy := make([]int, len(subset))
            copy(cpy, subset)
            ans = append(ans, cpy)
            return
        }

        subset = append(subset, nums[i])
        dfs(i + 1)

        subset = subset[:len(subset)-1]
        dfs(i + 1)
    }

    dfs(0)

    return ans
}
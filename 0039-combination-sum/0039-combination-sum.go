func combinationSum(candidates []int, target int) [][]int {
    ans := [][]int{}
    subset := []int{}

    var dfs func(i int, cur []int, total int)
    dfs = func(i int, cur []int, total int) {
        if total == target {
            temp := make([]int, len(subset))
            copy(temp, subset)
            ans = append(ans, temp)
            return
        }

        if i >= len(candidates) || total > target {
            return
        }

        subset = append(subset, candidates[i])
        dfs(i, subset, total + candidates[i])
        subset = subset[:len(subset)-1]
        dfs(i + 1, subset, total)
    }

    dfs(0, subset, 0)

    return ans
}
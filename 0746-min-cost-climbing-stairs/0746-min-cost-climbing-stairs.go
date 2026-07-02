func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    ans := make([]int, n)
    ans[0] = cost[0]
    ans[1] = cost[1]
    for i := 2; i < n; i++ {
        ans[i] = min(ans[i-1], ans[i-2]) + cost[i]
    }

    return min(ans[n-1], ans[n-2])
}
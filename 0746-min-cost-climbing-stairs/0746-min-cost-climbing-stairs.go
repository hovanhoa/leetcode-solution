func minCostClimbingStairs(cost []int) int {
    cost = append(cost, 0)
    for i := 2; i < len(cost); i++ {
        cost[i] += min(cost[i-2], cost[i-1])
    }

    return cost[len(cost)-1]
}
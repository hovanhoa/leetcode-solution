func equalPairs(grid [][]int) int {
    n := len(grid)

    rows := make(map[string]int)
    for i := 0; i < n; i++ {
        key := fmt.Sprint(grid[i])
        rows[key]++
    }

    ans := 0
    for j := 0; j < n; j++ {
        col := make([]int, n)
        for i := 0; i < n; i++ {
            col[i] = grid[i][j]
        }

        ans += rows[fmt.Sprint(col)]
    }

    return ans
}

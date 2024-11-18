func shiftGrid(grid [][]int, k int) [][]int {
    m, n := len(grid[0]), len(grid)
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            x := (k / m) % n
            y := k % m
            k++

            ans[x][y] = grid[i][j] 
        }
    }

    return ans
}
func numIslands(grid [][]byte) int {
    ans := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == '1' {
                ans += 1
                dfs(grid, i, j)
            }
        }
    }
    
    return ans
}

func dfs(grid [][]byte, row int, col int) {
    if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
        return
    }

    grid[row][col] = '0'
    dfs(grid, row + 1, col)
    dfs(grid, row - 1, col)
    dfs(grid, row, col + 1)
    dfs(grid, row, col - 1)
}
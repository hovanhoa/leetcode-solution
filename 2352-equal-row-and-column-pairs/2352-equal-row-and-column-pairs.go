func equalPairs(grid [][]int) int {
    n := len(grid)
    revertedGrid := make([][]int, n)
    for i := range revertedGrid {
        revertedGrid[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            revertedGrid[j][i] = grid[i][j]
        }
    }

    res := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if reflect.DeepEqual(grid[i], revertedGrid[j]) {
                res++
            }
        }
    }

    return res
}
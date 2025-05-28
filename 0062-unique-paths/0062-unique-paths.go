func uniquePaths(m int, n int) int {
    arr := make([][]int, m)
    for i := range arr {
        arr[i] = make([]int, n)
    }
    
    for i := range arr {
        arr[i][0] = 1
    }

    for i := range arr[0] {
        arr[0][i] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            arr[i][j] = arr[i-1][j] + arr[i][j-1]
        }
    }

    return arr[m-1][n-1]
}
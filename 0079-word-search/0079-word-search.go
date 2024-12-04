func exist(board [][]byte, word string) bool {
    mMap := map[[2]int]bool{}
    row, col := len(board), len(board[0])

    var dfs func(i, j, l int) bool
    dfs = func(i, j, l int) bool {
        if l == len(word) {
            return true
        }

        if i < 0 || j < 0 || i >= row || j >= col || board[i][j] != word[l] || mMap[[2]int{i, j}] {
            return false
        }

        mMap[[2]int{i, j}] = true
        res := dfs(i + 1, j, l + 1) || dfs(i - 1, j, l + 1) || dfs(i, j + 1, l + 1) || dfs(i, j - 1, l + 1)

        mMap[[2]int{i, j}] = false
        return res
    }

    for i := range board {
        for j := range board[0] {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}
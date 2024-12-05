func exist(board [][]byte, word string) bool {
    lenRow, lenCol := len(board), len(board[0])
    mMap := map[[2]int]bool{}
    
    var dfs func(i, j, l int) bool
    dfs = func(i, j, l int) bool {
        if l == len(word) {
            return true
        }

        if i < 0 || j < 0 || i == lenRow || j == lenCol || mMap[[2]int{i, j}]|| board[i][j] != word[l] {
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
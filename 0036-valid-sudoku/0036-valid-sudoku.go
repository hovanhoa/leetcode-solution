func isValidSudoku(board [][]byte) bool {
    cols := [9][9]bool{}
    rows := [9][9]bool{}
    squares := [3][3][9]bool{}
    for i := range board {
        for j := range board[i] {
            cell := board[i][j]
            if cell == '.' {
                continue
            }

            digit := int(cell - '0') - 1
            if cols[i][digit] || rows[j][digit] || squares[i/3][j/3][digit] {
                return false
            }

            cols[i][digit] = true
            rows[j][digit] = true
            squares[i/3][j/3][digit] = true
        }
    }

    return true
}
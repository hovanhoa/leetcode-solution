func setZeroes(matrix [][]int)  {
    mRow, mCol := map[int]bool{}, map[int]bool{}
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == 0 {
                mRow[i] = true
                mCol[j] = true
            }
        }
    }

    for i := range matrix {
        for j := range matrix[i] {
            if mRow[i] || mCol[j] {
                matrix[i][j] = 0
            }
        }
    }
}
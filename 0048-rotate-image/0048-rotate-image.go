func rotate(matrix [][]int)  {
    l, r := 0, len(matrix) - 1
    for l < r {
        for i := 0; i < r - l; i++ {
            top, bot := l, r
            topLeft := matrix[top][l + i]

            matrix[top][l + i] = matrix[bot-i][l]
            matrix[bot-i][l] = matrix[bot][r-i]
            matrix[bot][r-i] = matrix[top+i][r]
            matrix[top+i][r] = topLeft
        }

        l += 1
        r -= 1
    }
}

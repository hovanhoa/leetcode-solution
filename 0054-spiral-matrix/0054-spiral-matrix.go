func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    ans := make([]int, 0, len(matrix) * len(matrix[0]))
    left, right := 0, len(matrix[0]) - 1
    top, bottom := 0, len(matrix) - 1
    for {
        for i := left; i <= right; i++ {
            ans = append(ans, matrix[top][i])
        }
        top += 1
        if top > bottom {
            break
        }

        for i := top; i <= bottom; i++ {
            ans = append(ans, matrix[i][right])
        }
        right -= 1
        if right < left {
            break
        }

        for i := right; i >= left; i-- {
            ans = append(ans, matrix[bottom][i])
        } 
        bottom -= 1
        if bottom < top {
            break
        }

        for i := bottom; i >= top; i-- {
            ans = append(ans, matrix[i][left])
        } 
        left += 1
        if left > right {
            break
        }
    }

    return ans
}
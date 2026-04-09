func largestRectangleArea(heights []int) int {
    stack := []int{}
    res := 0
    n := len(heights)

    for i := 0; i <= n; i++ {
        newHeight := 0
        if i < n {
            newHeight = heights[i]
        }

        for len(stack) > 0 && heights[stack[len(stack)-1]] > newHeight {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            res = max(res, height*width)
        }

        stack = append(stack, i)
    }

    return res
}
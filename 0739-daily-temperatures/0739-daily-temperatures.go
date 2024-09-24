func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    stack := make([]int, 0)
    for i, v := range temperatures {
        for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
            ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }

        stack = append(stack, i)
    }

    return ans
}
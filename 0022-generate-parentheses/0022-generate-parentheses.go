func generateParenthesis(n int) []string {
    ans := []string{}
    
    var backtracking func(cur string, open int, close int)
    backtracking = func(cur string, open int, close int) {
        if len(cur) == 2 * n {
            ans = append(ans, cur)
            return
        }

        if open < n {
            backtracking(cur + "(", open + 1, close)
        }

        if close < open && close < n {
            backtracking(cur + ")", open, close + 1)
        }
    }

    backtracking("", 0, 0)
    return ans
}
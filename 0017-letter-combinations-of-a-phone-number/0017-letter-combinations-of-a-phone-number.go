func letterCombinations(digits string) []string {
    ans := []string{}
    mPhone := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    var backtracking func(cur string, i int)
    backtracking = func(cur string, i int) {
        if len(cur) == len(digits) {
            ans = append(ans, cur)
            return
        }

        for _, v := range mPhone[digits[i]] {
            backtracking(cur + string(v), i + 1)
        }
    }

    if len(digits) > 0 {
        backtracking("", 0)
    }

    return ans
}
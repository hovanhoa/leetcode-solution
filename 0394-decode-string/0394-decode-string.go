func decodeString(s string) string {
    nums, codes := []int{}, []string{}
    num, code := "", ""
    for _, c := range s {
        if unicode.IsNumber(c) {
            num += string(c)
        } else if unicode.IsLetter(c) {
            code += string(c)
        } else if c == '[' {
            n, _ := strconv.Atoi(num)
            nums, codes = append(nums, n), append(codes, code)
            num, code = "", ""
        } else if c == ']' {
            n, c := nums[len(nums)-1], codes[len(codes)-1]
            code = c + strings.Repeat(code, n)
            nums, codes = nums[:len(nums)-1], codes[:len(codes)-1]
        }
    }
    return code
}
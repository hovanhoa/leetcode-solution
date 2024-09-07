func removeStars(s string) string {
    var stack []byte
    for _, v := range s {
        if v == '*' {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, byte(v))
        }
    } 

    return string(stack)
}
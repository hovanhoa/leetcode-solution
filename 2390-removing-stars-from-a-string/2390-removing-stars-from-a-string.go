func removeStars(s string) string {
    stack := []byte{}
    for _, v := range s {
        if byte(v) == '*' {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, byte(v))
        }
    }

    return string(stack)
}

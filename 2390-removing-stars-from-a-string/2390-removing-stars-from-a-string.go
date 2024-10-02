func removeStars(s string) string {
    ans := []byte{}
    for _, v := range []byte(s) {
        if v == '*' {
            ans = ans[:len(ans)-1]
        } else {
            ans = append(ans, v)
        }
    }

    return string(ans)
}
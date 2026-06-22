import "strconv"

func compress(chars []byte) int {
    w := 0
    r := 0
    for r < len(chars) {
        c, cnt := chars[r], 0
        for r < len(chars) && c == chars[r] {
            r += 1
            cnt += 1
        }

        chars[w] = c
        w += 1
        if cnt > 1 {
            for _, d := range strconv.Itoa(cnt) {
                chars[w] = byte(d)
                w += 1
            }
        }
    }
    return w
}

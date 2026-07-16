func numJewelsInStones(jewels string, stones string) int {
    m := map[byte]bool{}
    for _, v := range jewels {
        m[byte(v)] = true
    }

    cnt := 0
    for _, v := range stones {
        if m[byte(v)] {
            cnt += 1
        }
    }

    return cnt
}
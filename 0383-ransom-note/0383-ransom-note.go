func canConstruct(ransomNote string, magazine string) bool {
    m := map[byte]int{}
    for _, v := range magazine {
        m[byte(v)] += 1
    }

    for _, v := range ransomNote {
        if m[byte(v)] > 0 {
            m[byte(v)] -= 1
        } else {
            return false
        }
    }

    return true
}
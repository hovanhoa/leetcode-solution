func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }    

    m1, m2 := map[byte]int{}, map[byte]int{}
    for _, v := range word1 {
        m1[byte(v)] += 1
    }

    for _, v := range word2 {
        m2[byte(v)] += 1
    }

    fmt.Println(m1)
    fmt.Println(m2)

    return isSameMap(m1, m2) && isSameMap(m2, m1)
}

func isSameMap(m1, m2 map[byte]int) bool {
    if len(m1) != len(m2) {
        return false
    }

    reverseMap := map[int]int{}
    for _, v := range m2 {
        reverseMap[v] += 1
    }

    for k, v := range m1 {
        if _, ok := m2[k]; !ok {
            return false
        }

        if reverseMap[v] > 0 {
            reverseMap[v] -= 1
        } else {
            return false
        }
    }

    return true
}

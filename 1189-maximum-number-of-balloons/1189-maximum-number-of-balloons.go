func maxNumberOfBalloons(text string) int {
    m := map[byte]int{}
    for _, v := range text {
        m[byte(v)] += 1
    }

    ans := 0
    for {
        for _, v := range "balloon" {
            if m[byte(v)] > 0 {
                m[byte(v)] -= 1
            } else {
                return ans
            }
        }
        
        ans += 1
    }
}
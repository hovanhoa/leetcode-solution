func wordPattern(pattern string, s string) bool {
    sArr := strings.Split(s, " ")
    if len(sArr) != len(pattern) {
        return false
    }

    m := map[byte]string{}
    mInvert := map[string]byte{}
    for i := 0; i < len(pattern); i++ {
        if m[pattern[i]] != "" && m[pattern[i]] != sArr[i] {
            return false
        }

        if mInvert[sArr[i]] != byte(0) && mInvert[sArr[i]] != pattern[i] {
            return false
        }

        m[pattern[i]] = sArr[i]
        mInvert[sArr[i]] = pattern[i]
    } 

    return true
}
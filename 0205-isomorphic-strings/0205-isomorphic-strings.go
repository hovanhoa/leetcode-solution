func isIsomorphic(s string, t string) bool {
    mapST, mapTS := map[byte]byte{}, map[byte]byte{}
    zero := byte(0)
    for i := 0; i < len(s); i++ {
        st, ts := s[i], t[i]

        if mapST[st] != zero && mapST[st] != ts {
            return false
        }

        if mapTS[ts] != zero && mapTS[ts] != st {
            return false
        }

        mapST[st] = ts
        mapTS[ts] = st
    }

    return true
}
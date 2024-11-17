func reverseString(s []byte)  {
    l, r := 0, len(s) - 1
    for l <= r {
        s[l], s[r] = s[r], s[l]
        l, r = l + 1, r - 1
    }
}
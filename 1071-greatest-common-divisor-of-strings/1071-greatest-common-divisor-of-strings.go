func gcdOfStrings(str1 string, str2 string) string {
    if str1 == str2 {
        return str1
    }

    if str1 == "" || str2 == "" {
        return ""
    }

    if len(str2) > len(str1) {
        return gcdOfStrings(str2, str1)
    }

    return gcdOfStrings(str2, str1[len(str2):])
}
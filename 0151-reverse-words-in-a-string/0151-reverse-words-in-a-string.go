import "strings"

func reverseWords(s string) string {
    arr := strings.Split(s, " ")
    filtered := []string{}
    for i := len(arr) - 1; i >= 0; i -= 1 {
        if arr[i] != "" {
            filtered = append(filtered, arr[i])
        }
    }

    return strings.Join(filtered, " ")
}
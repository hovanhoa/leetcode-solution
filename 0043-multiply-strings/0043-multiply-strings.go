func multiply(num1 string, num2 string) string {
    n1, n2 := num1, num2
    if n1 == "0" || n2 == "0" {
        return "0"
    }

    ans := make([]byte, len(n1) + len(n2))
    for i := len(n1) - 1; i >= 0; i-- {
        for j := len(n2) - 1; j >= 0; j-- {
            v1 := n1[i] - '0'
            v2 := n2[j] - '0'
            ans[i+j+1] += (v1*v2)
            if ans[i+j+1] >= 10 {
                ans[i+j] += ans[i+j+1]/10
                ans[i+j+1] = ans[i+j+1]%10 
            }
        }
    }

    if ans[0] == 0 {
        ans = ans[1:]
    }

    for i := range ans {
        ans[i] += '0'
    }

    return string(ans)
}
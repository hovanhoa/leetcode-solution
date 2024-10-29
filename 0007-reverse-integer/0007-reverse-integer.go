func reverse(x int) int {
    x32 := int32(x)
    ans := int32(0)
    for x32 != 0 {
        d := x32 % 10
        x32 /= 10

        newAns := ans * 10 + d
        if (newAns - d) / 10 != ans {
            return 0
        }

        ans = newAns
    }

    return int(ans)
}
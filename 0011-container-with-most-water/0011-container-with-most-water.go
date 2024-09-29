func maxArea(height []int) int {
    ans := 0
    l, r := 0, len(height) - 1
    for l < r {
        cur := (r - l) * min(height[l], height[r])
        if ans < cur {
            ans = cur
        }

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
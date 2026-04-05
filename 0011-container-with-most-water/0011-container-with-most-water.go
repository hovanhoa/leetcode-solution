func maxArea(height []int) int {
    res := 0
    l, r := 0, len(height) - 1
    for l < r {
        res = max(res, (r - l) * min(height[l], height[r]))

        if height[l] < height[r] {
            l += 1
        } else {
            r -= 1
        }
    }

    return res
}
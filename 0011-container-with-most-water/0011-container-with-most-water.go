func maxArea(height []int) int {
    maxArea := 0
    l, r := 0, len(height) - 1
    for l < r {
        newArea := (r - l) * min(height[l], height[r])
        maxArea = max(maxArea, newArea)

        if height[l] < height[r] {
            l += 1
        } else {
            r -= 1
        }
    }

    return maxArea
}
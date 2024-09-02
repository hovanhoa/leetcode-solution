func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    maxArea := r * min(height[0], height[r])

    for l < r {
        if height[l] <= height[r] {
            l++
        } else {
            r--
        }

        if l < r && (r - l) * min(height[r], height[l]) > maxArea {
            maxArea = (r - l) * min(height[r], height[l])
        }
    }

    return maxArea
}
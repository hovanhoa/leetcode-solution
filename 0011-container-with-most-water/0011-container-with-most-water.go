func maxArea(height []int) int {
    ans := 0
    left, right := 0, len(height) - 1
    for left < right {
        ans = max(ans, min(height[left], height[right])*(right-left))
        if height[left] > height[right] {
            right -= 1
        } else {
            left += 1
        }
    }

    return ans
}
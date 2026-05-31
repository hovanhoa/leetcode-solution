func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers) - 1
    s := numbers[left] + numbers[right]
    for left < right {
        if s < target {
            left += 1
            s += numbers[left] - numbers[left-1]
        } else if s > target {
            right -= 1
            s += numbers[right] - numbers[right+1]
        } else {
            return []int{left+1, right+1}
        }
    }

    return []int{}
}
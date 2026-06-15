func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1
    s := numbers[l] + numbers[r]
    for l < r {
        if s < target {
            l += 1
            s += numbers[l] - numbers[l-1]
        } else if s > target {
            r -= 1
            s += numbers[r] - numbers[r+1]
        } else {
            return []int{l+1, r+1}
        }
    }

    return []int{}
}

func findDisappearedNumbers(nums []int) []int {
    // Step 1: Mark visited indices as negative
    for _, num := range nums {
        index := abs(num) - 1
        if nums[index] > 0 {
            nums[index] = -nums[index]
        }
    }
    
    // Step 2: Identify missing numbers by positive elements
    var result []int
    for i, num := range nums {
        if num > 0 {
            result = append(result, i+1)
        }
    }
    
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
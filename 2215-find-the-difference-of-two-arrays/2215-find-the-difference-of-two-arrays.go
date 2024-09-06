func findDifference(nums1 []int, nums2 []int) [][]int {
    result1 := []int{}
    for _, val := range nums1 {
        if !slices.Contains(nums2, val) && !slices.Contains(result1, val) {
            result1 = append(result1, val)
        }
    }

    result2 := []int{}
    for _, val := range nums2 {
        if !slices.Contains(nums1, val) && !slices.Contains(result2, val) {
            result2 = append(result2, val)
        }
    }

    return [][]int{result1, result2}
}
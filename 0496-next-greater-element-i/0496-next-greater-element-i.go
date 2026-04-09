func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    stack := []int{}
    for _, v := range nums2 {
        for len(stack) > 0 && v > stack[len(stack)-1] {
            m[stack[len(stack)-1]] = v
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, v)
    }

    res := make([]int, len(nums1))
    for i, v := range nums1 {
        res[i] = -1
        if val, ok := m[v]; ok {
            res[i] = val
        }
    }

    return res
}
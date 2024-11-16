func nextGreaterElement(nums1 []int, nums2 []int) []int {
    ans := make([]int, len(nums1))
    m := map[int]int{}
    for i, v := range nums1 {
        m[v] = i
    }

    for i, v := range nums2 {
        if _, ok := m[v]; !ok {
            continue   
        } else {
            idx := m[v]
            ans[idx] = -1
            for j := i + 1; j < len(nums2); j++ {
                if nums2[j] > nums2[i] {
                    ans[idx] = nums2[j]
                    break
                }
            }
        }
    }

    return ans

    
}
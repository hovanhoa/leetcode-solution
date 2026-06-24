func findDifference(nums1 []int, nums2 []int) [][]int {
    m1, m2 := map[int]bool{}, map[int]bool{}
    for _, v := range nums1 {
        m1[v] = true
    }

    for _, v := range nums2 {
        m2[v] = true
    }

    ans1 := []int{}
    for _, v := range nums1 {
        if !m2[v] {
            ans1 = append(ans1, v)
        }

        m2[v] = true
    }
    
    ans2 := []int{}
    for _, v := range nums2 {
        if !m1[v] {
            ans2 = append(ans2, v)
        }

        m1[v] = true
    }

    return [][]int{ans1, ans2}
}
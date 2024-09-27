func findDifference(nums1 []int, nums2 []int) [][]int {
    s1Map, s2Map := map[int]bool{}, map[int]bool{}
    for _, v := range nums1 {
        s1Map[v] = true
    }

    for _, v := range nums2 {
        s2Map[v] = true
    }

    ans := make([][]int, 2)
    for _, v := range nums1 {
        if !s2Map[v] {
            ans[0] = append(ans[0], v)
            s2Map[v] = true
        }
    }

    for _, v := range nums2 {
        if !s1Map[v] {
            ans[1] = append(ans[1], v)
            s1Map[v] = true
        }
    }

    return ans
}
func findDifference(nums1 []int, nums2 []int) [][]int {
    s1Map := make(map[int]bool)
    s2Map := make(map[int]bool)

    for _, v := range nums1 {
        s1Map[v] = true
    }

    for _, v := range nums2 {
        s2Map[v] = true
    }

    res := make([][]int, 2)
    for k := range s1Map {
        if !s2Map[k] {
            res[0] = append(res[0], k)
        }
    }

    for k := range s2Map {
        if !s1Map[k] {
            res[1] = append(res[1], k)
        }
    }

    return res
}
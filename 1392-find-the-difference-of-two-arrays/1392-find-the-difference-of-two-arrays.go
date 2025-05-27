func findDifference(nums1 []int, nums2 []int) [][]int {
    m1, m2 := map[int]bool{}, map[int]bool{}
    for _, v := range nums1 {
        m1[v] = true
    }

    for _, v := range nums2 {
        m2[v] = true
    }

    ans := [][]int{}

    arr := []int{}
    for _, v := range nums1 {
        if !m2[v] {
            arr = append(arr, v)
        }

        m2[v] = true
    }
    ans = append(ans, arr)

    arr = []int{}
    for _, v := range nums2 {
        if !m1[v] {
            arr = append(arr, v)
        }

        m1[v] = true
    }
    ans = append(ans, arr)

    return ans

}
func findDisappearedNumbers(nums []int) []int {
    seen, ans := map[int]bool{}, []int{}
    for _, v := range nums {
        seen[v] = true
    }

    for i := 1; i <= len(nums); i++ {
        if !seen[i] {
            ans = append(ans, i)
        }
    }

    return ans
}
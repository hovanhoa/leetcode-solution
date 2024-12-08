func permuteUnique(nums []int) [][]int {
    ans := [][]int{}
    perm := []int{}
    count := map[int]int{}
    for _, v := range nums {
        count[v] += 1
    }

    var dfs func()
    dfs = func() {
        if len(perm) == len(nums) {
            tmp := make([]int, len(perm))
            copy(tmp, perm)
            ans = append(ans, tmp)
            return
        }

        for k := range count {
            if count[k] > 0 {
                perm = append(perm, k)
                count[k] -= 1

                dfs()

                count[k] += 1
                perm = perm[:len(perm)-1]
            }
        } 
    }

    dfs()
    return ans
}
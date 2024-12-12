func findOrder(numCourses int, prerequisites [][]int) []int {
    preMap := map[int][]int{}
    for _, v := range prerequisites {
        preMap[v[0]] = append(preMap[v[0]], v[1])
    }

    visit, cycle := map[int]bool{}, map[int]bool{}
    ans := []int{}
    var dfs func(cur int) bool
    dfs = func(cur int) bool {
        if cycle[cur] {
            return false
        }

        if visit[cur] {
            return true
        }

        cycle[cur] = true
        for _, v := range preMap[cur] {
            if !dfs(v) {
                return false
            }
        }

        cycle[cur] = false
        visit[cur] = true
        ans = append(ans, cur)
        return true
    }

    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return []int{}
        }
    }

    return ans
}
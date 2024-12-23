func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    ans := [][]int{}
    cur := []int{intervals[0][0], intervals[0][1]}
    for i := range intervals {
        if intervals[i][0] <= cur[1] {
            cur[1] = max(intervals[i][1], cur[1])
            continue
        } else {
            ans = append(ans, cur)
            cur = intervals[i]
        }
    }

    ans = append(ans, cur)

    return ans
}
func insert(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
    return merge(intervals)
}

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    ans := [][]int{}
    cur := []int{intervals[0][0], intervals[0][1]}
    for i := range intervals {
        if cur[1] >= intervals[i][0] {
            cur[1] = max(cur[1], intervals[i][1])
            continue
        } else {
            ans = append(ans, cur)
            cur = intervals[i]
        }
    }

    ans = append(ans, cur)

    return ans
}
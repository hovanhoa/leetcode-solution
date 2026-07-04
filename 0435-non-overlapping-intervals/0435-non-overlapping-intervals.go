func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[1] - b[1]
    })

    cur := intervals[0][1]
    cnt := 1
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] >= cur {
            cnt += 1
            cur = intervals[i][1]
        }
    }

    return len(intervals) - cnt
}
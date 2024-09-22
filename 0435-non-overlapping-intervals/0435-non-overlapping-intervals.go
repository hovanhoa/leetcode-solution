func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) <= 1 {
        return 0
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    cur := intervals[0]
    count := 1
    for _, v := range intervals[1:] {
        if v[0] >= cur[1] {
            cur = v
            count++
        }
    }

    return len(intervals) - count
}
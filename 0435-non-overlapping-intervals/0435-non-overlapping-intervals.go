func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    minE := math.MaxInt
    count := 1
    for _, v := range intervals {
        if v[0] >= minE {
            count++
            minE = v[1]
        } else {
            if v[1] < minE {
                minE = v[1]
            }
        }
    }

    return len(intervals) - count
}
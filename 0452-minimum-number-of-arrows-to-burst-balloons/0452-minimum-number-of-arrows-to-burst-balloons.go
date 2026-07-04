func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int {
        return a[1] - b[1]
    })

    cur := points[0][1]
    cnt := 1
    for i := 1; i < len(points); i++ {
        if points[i][0] > cur {
            cnt += 1
            cur = points[i][1]
        }
        
    }

    return cnt
}
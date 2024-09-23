func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(x, y []int) int {
        return x[1] - y[1]
    })

    cur := points[0][1]
    ans := 1
    for _, v := range points {
        if v[0] > cur {
            cur = v[1]
            ans++
        }
    }

    return ans
}
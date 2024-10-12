func minReorder(n int, connections [][]int) int {
    arr := make([][]int, n)
    for i := range connections {
        f, s := connections[i][0], connections[i][1]
        arr[f] = append(arr[f], s)
        arr[s] = append(arr[s], -f)
    }


    ans := 0
    visited := make([]bool, n)
    var dfs func(num int)
    dfs = func(num int) {
        visited[num] = true
        for i := range arr[num] {
            if !visited[abs(arr[num][i])] {
                dfs(abs(arr[num][i]))
                if arr[num][i] > 0 {
                    ans += 1
                }
            }
        }
    }

    for i := range arr {
        dfs(i)
    }

    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}
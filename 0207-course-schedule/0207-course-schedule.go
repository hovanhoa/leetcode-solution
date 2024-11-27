func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for i := 0; i < numCourses; i++ {
        graph[i] = []int{}
    }

    for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

    visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && hasCycle(graph, visited, i) {
			return false
		}
	}

	return true
}

func hasCycle(graph [][]int, visited[]int, node int) bool {
    if visited[node] == 1 {
        return true
    }

    if visited[node] == -1 {
        return false
    }

    visited[node] = 1
    for _, neighbor := range graph[node] {
        if hasCycle(graph, visited, neighbor) {
			return true
		}
    }

    visited[node] = -1
    return false
}
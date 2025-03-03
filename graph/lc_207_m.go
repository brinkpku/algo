package graph

// dfs 环检测
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	track := make([]bool, numCourses)
	hasCycle := false
	var dfs func(i int)
	dfs = func(i int) {
		if track[i] {
			hasCycle = true
			return
		}
		if hasCycle || visited[i] {
			return
		}
		track[i] = true
		visited[i] = true
		for _, n := range graph[i] {
			dfs(n)
		}
		track[i] = false
	}
	for i := range numCourses {
		dfs(i)
	}
	return !hasCycle
}

func buildGraph(nums int, relation [][]int) [][]int {
	graph := make([][]int, nums)
	for i := range relation {
		graph[relation[i][0]] = append(graph[relation[i][0]], relation[i][1])
	}
	return graph
}

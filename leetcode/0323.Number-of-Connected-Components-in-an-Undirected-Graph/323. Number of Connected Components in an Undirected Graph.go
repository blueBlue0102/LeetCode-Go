package leetcode

func NumberofConnectedComponentsinanUndirectedGraph(n int, edges [][]int) int {
	edgeMap := map[int][]int{}
	for i := range edges {
		node1 := edges[i][0]
		node2 := edges[i][1]
		if edgeMap[node1] == nil {
			edgeMap[node1] = make([]int, 0)
		}
		if edgeMap[node2] == nil {
			edgeMap[node2] = make([]int, 0)
		}

		edgeMap[node1] = append(edgeMap[node1], node2)
		edgeMap[node2] = append(edgeMap[node2], node1)
	}

	visit := make([]bool, n)
	var dfs func(node int)
	dfs = func(node int) {
		if visit[node] {
			return
		}
		visit[node] = true
		for _, neighbor := range edgeMap[node] {
			dfs(neighbor)
		}
	}

	count := 0
	for node := 0; node < n; node++ {
		if !visit[node] {
			count++
			dfs(node)
		}
	}

	return count
}

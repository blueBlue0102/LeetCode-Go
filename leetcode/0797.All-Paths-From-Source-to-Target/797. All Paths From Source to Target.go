package leetcode

// 建立哪些 value 節點可以到達 key 節點的 map
// e.g. nodeMap[0] = [1, 2, 3] 表示節點 1, 2, 3 可以到達節點 0
var nodeMap map[int][]int

func AllPathsFromSourcetoTarget(graph [][]int) [][]int {
	// return Version_1(graph)
	return Version_2(graph)
}

func Enumerate_v1(res *[][]int, target int, path []int) {
	if target == 0 {
		*res = append(*res, append([]int{target}, path...))
		return
	}
	for _, node := range nodeMap[target] {
		Enumerate_v1(res, node, append([]int{target}, path...))
	}
}

func Version_1(graph [][]int) [][]int {
	result := [][]int{}

	nodeMap = map[int][]int{}
	for node, dests := range graph {
		for _, dest := range dests {
			if nodeMap[dest] == nil {
				nodeMap[dest] = []int{}
			}
			nodeMap[dest] = append(nodeMap[dest], node)
		}
	}

	Enumerate_v1(&result, len(graph)-1, []int{})

	return result
}

func Version_2(graph [][]int) [][]int {
	result := [][]int{}
	Enumerate_v2(&result, graph, 0, len(graph)-1, []int{})
	return result
}

func Enumerate_v2(res *[][]int, graph [][]int, curr int, target int, path []int) {
	if curr == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, append(tmp, target))
		return
	}
	for _, node := range graph[curr] {
		path = append(path, curr)
		Enumerate_v2(res, graph, node, target, path)
		path = path[:len(path)-1]
	}
}

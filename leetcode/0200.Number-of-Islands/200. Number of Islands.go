package leetcode

type UnionFind struct {
	parent    []int
	rank      []int
	rootCount int
}

// 變種的 Union Find，由於並不知道島嶼數量，所以並不會一開始就 assign 所有 node，而是後續再增加
func NewUnionFind(size int) *UnionFind {
	unionFind := &UnionFind{
		parent:    make([]int, 0, size),
		rank:      make([]int, 0, size),
		rootCount: 0, // 記錄現在 parent 中有多少個 root
	}
	return unionFind
}

func (uf *UnionFind) AddMember(node int) {
	uf.parent = append(uf.parent, node)
	uf.rank = append(uf.rank, 0)
	uf.rootCount++
}

func (uf *UnionFind) Find(node int) int {
	if uf.parent[node] != node {
		// path compression
		uf.parent[node] = uf.Find(uf.parent[node])
	}
	return uf.parent[node]
}

func (uf *UnionFind) Union(n1 int, n2 int) {
	rootN1 := uf.Find(n1)
	rootN2 := uf.Find(n2)

	// union by rank
	if rootN1 != rootN2 {
		uf.rootCount--
		if uf.rank[rootN1] > uf.rank[rootN2] {
			uf.parent[rootN2] = rootN1
		} else if uf.rank[rootN1] < uf.rank[rootN2] {
			uf.parent[rootN1] = rootN2
		} else {
			uf.parent[rootN2] = rootN1
			uf.rank[rootN1]++
		}
	}
}

func NumberofIslands(grid [][]byte) int {
	// island 的號碼，從 1 號開始
	islandNumber := 0
	// 根據 island 號碼記錄的 map
	islandMap := make([][]int, len(grid))
	for i := range islandMap {
		islandMap[i] = make([]int, len(grid[0]))
	}
	unionFind := NewUnionFind(len(grid) * len(grid[0]))

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				// 上方座標的島嶼號碼
				var upperIslandNum int
				// 左方座標的島嶼號碼
				var leftIslandNum int
				if row > 0 && islandMap[row-1][col] != 0 {
					upperIslandNum = islandMap[row-1][col]
				}
				if col > 0 && islandMap[row][col-1] != 0 {
					leftIslandNum = islandMap[row][col-1]
				}

				if upperIslandNum == 0 && leftIslandNum == 0 {
					// maybe a new island
					islandNumber++
					islandMap[row][col] = islandNumber
					unionFind.AddMember(islandNumber - 1) // union find 的號碼是從 0 開始，所以要減 1
				} else if upperIslandNum != 0 && leftIslandNum == 0 {
					islandMap[row][col] = upperIslandNum
				} else if upperIslandNum == 0 && leftIslandNum != 0 {
					islandMap[row][col] = leftIslandNum
				} else {
					// upperGrid != 0 && leftGrid != 0
					// 要進行 union
					islandMap[row][col] = leftIslandNum
					unionFind.Union(upperIslandNum-1, leftIslandNum-1) // union find 的號碼是從 0 開始，所以要減 1
				}
			}
		}
	}

	return unionFind.rootCount
}

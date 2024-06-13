package structures

// 經過 union by rank 以及 path compression 優化後的 Quick Union 實作

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	unionFind := &UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range unionFind.parent {
		unionFind.parent[i] = i
	}
	return unionFind
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

func (uf *UnionFind) IsConnected(n1 int, n2 int) bool {
	return uf.Find(n1) == uf.Find(n2)
}

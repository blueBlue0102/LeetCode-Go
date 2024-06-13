package structures

import (
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	for i := 0; i < 10; i++ {
		if uf.parent[i] != i {
			t.Errorf("expected parent[%d] == %d, got %d", i, i, uf.parent[i])
		}
		if uf.rank[i] != 0 {
			t.Errorf("expected rank[%d] == 0, got %d", i, uf.rank[i])
		}
	}
}

func TestFind(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Union(0, 1)
	uf.Union(1, 2)

	if uf.Find(0) != uf.Find(1) {
		t.Errorf("expected Find(0) == Find(1), got %d and %d", uf.Find(0), uf.Find(1))
	}
	if uf.Find(1) != uf.Find(2) {
		t.Errorf("expected Find(1) == Find(2), got %d and %d", uf.Find(1), uf.Find(2))
	}
	if uf.Find(0) == uf.Find(3) {
		t.Errorf("expected Find(0) != Find(3), got %d and %d", uf.Find(0), uf.Find(3))
	}
}

func TestUnion(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Union(0, 1)
	if uf.Find(0) != uf.Find(1) {
		t.Errorf("expected Find(0) == Find(1), got %d and %d", uf.Find(0), uf.Find(1))
	}

	uf.Union(1, 2)
	if uf.Find(0) != uf.Find(2) {
		t.Errorf("expected Find(0) == Find(2), got %d and %d", uf.Find(0), uf.Find(2))
	}

	uf.Union(3, 4)
	if uf.Find(3) != uf.Find(4) {
		t.Errorf("expected Find(3) == Find(4), got %d and %d", uf.Find(3), uf.Find(4))
	}
	if uf.Find(0) == uf.Find(3) {
		t.Errorf("expected Find(0) != Find(3), got %d and %d", uf.Find(0), uf.Find(3))
	}

	uf.Union(2, 4)
	if uf.Find(0) != uf.Find(4) {
		t.Errorf("expected Find(0) == Find(4), got %d and %d", uf.Find(0), uf.Find(4))
	}
}

func TestPathCompression(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 4)

	if uf.Find(0) != uf.Find(4) {
		t.Errorf("expected Find(0) == Find(4), got %d and %d", uf.Find(0), uf.Find(4))
	}

	// Check if path compression worked
	if uf.parent[1] != uf.Find(1) {
		t.Errorf("expected parent[1] == Find(1), got %d and %d", uf.parent[1], uf.Find(1))
	}
	if uf.parent[2] != uf.Find(2) {
		t.Errorf("expected parent[2] == Find(2), got %d and %d", uf.parent[2], uf.Find(2))
	}
	if uf.parent[3] != uf.Find(3) {
		t.Errorf("expected parent[3] == Find(3), got %d and %d", uf.parent[3], uf.Find(3))
	}
	if uf.parent[4] != uf.Find(4) {
		t.Errorf("expected parent[4] == Find(4), got %d and %d", uf.parent[4], uf.Find(4))
	}
}

func TestUnionByRank(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)

	initialRank := uf.rank[uf.Find(0)]

	uf.Union(0, 2)
	if uf.Find(0) != uf.Find(2) {
		t.Errorf("expected Find(0) == Find(2), got %d and %d", uf.Find(0), uf.Find(2))
	}
	if uf.rank[uf.Find(0)] != initialRank+1 {
		t.Errorf("expected rank == %d, got %d", initialRank+1, uf.rank[uf.Find(0)])
	}

	uf.Union(0, 4)
	if uf.Find(0) != uf.Find(4) {
		t.Errorf("expected Find(0) == Find(4), got %d and %d", uf.Find(0), uf.Find(4))
	}
	if uf.rank[uf.Find(0)] != initialRank+1 {
		t.Errorf("expected rank == %d, got %d", initialRank+1, uf.rank[uf.Find(0)])
	}
}

func TestDisconnectedSets(t *testing.T) {
	uf := NewUnionFind(10)
	if uf.Find(0) == uf.Find(9) {
		t.Errorf("expected Find(0) != Find(9), got %d and %d", uf.Find(0), uf.Find(9))
	}

	uf.Union(0, 9)
	if uf.Find(0) != uf.Find(9) {
		t.Errorf("expected Find(0) == Find(9), got %d and %d", uf.Find(0), uf.Find(9))
	}

	uf.Union(1, 2)
	if uf.Find(0) == uf.Find(1) {
		t.Errorf("expected Find(0) != Find(1), got %d and %d", uf.Find(0), uf.Find(1))
	}
	if uf.Find(1) != uf.Find(2) {
		t.Errorf("expected Find(1) == Find(2), got %d and %d", uf.Find(1), uf.Find(2))
	}
}

func TestUnionAll(t *testing.T) {
	uf := NewUnionFind(10)
	for i := 0; i < 9; i++ {
		uf.Union(i, i+1)
	}

	for i := 0; i < 10; i++ {
		if uf.Find(0) != uf.Find(i) {
			t.Errorf("expected Find(0) == Find(%d), got %d and %d", i, uf.Find(0), uf.Find(i))
		}
	}
}

func TestIsConnected(t *testing.T) {
	uf := NewUnionFind(10)
	if uf.IsConnected(0, 1) {
		t.Errorf("expected IsConnected(0, 1) == false, got true")
	}

	uf.Union(0, 1)
	if !uf.IsConnected(0, 1) {
		t.Errorf("expected IsConnected(0, 1) == true, got false")
	}

	uf.Union(2, 3)
	if uf.IsConnected(0, 2) {
		t.Errorf("expected IsConnected(0, 2) == false, got true")
	}
	if !uf.IsConnected(2, 3) {
		t.Errorf("expected IsConnected(2, 3) == true, got false")
	}

	uf.Union(1, 2)
	if !uf.IsConnected(0, 3) {
		t.Errorf("expected IsConnected(0, 3) == true, got false")
	}
}

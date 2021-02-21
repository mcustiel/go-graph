package graph

import "testing"

func createTestGraph() *Graph {
	root := NewNode(0)
	one := NewNode(1)
	two := NewNode(2)
	three := NewNode(3)
	four := NewNode(4)
	five := NewNode(5)
	six := NewNode(6)
	seven := NewNode(7)
	eight := NewNode(8)
	nine := NewNode(9)

	root.AddAdyacent(one)
	root.AddAdyacent(two)
	root.AddAdyacent(three)
	one.AddAdyacent(four)
	one.AddAdyacent(five)
	two.AddAdyacent(six)
	two.AddAdyacent(seven)
	three.AddAdyacent(eight)
	four.AddAdyacent(nine)

	return New(root)
}

func sliceEquals(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func TestWalk(t *testing.T) {
	graph := createTestGraph()
	list := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	cur := 0
	graph.Walk(func(node *Node) {
		list[cur] = node.Value().(int)
		cur++
	})
	expected := []int{0, 1, 4, 9, 5, 2, 6, 7, 3, 8}
	if !sliceEquals(expected[:], list[:]) {
		t.Errorf("Array is incorrect, got: %v, want: %v.", list, expected)
	}
}

func TestBfs(t *testing.T) {
	graph := createTestGraph()
	list := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	cur := 0
	graph.Bfs(func(node *Node) {
		list[cur] = node.Value().(int)
		cur++
	})
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !sliceEquals(expected[:], list[:]) {
		t.Errorf("Array is incorrect, got: %v, want: %v.", list, expected)
	}
}

func TestDfs(t *testing.T) {
	graph := createTestGraph()
	list := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	cur := 0
	graph.Dfs(func(node *Node) {
		list[cur] = node.Value().(int)
		cur++
	})
	expected := []int{0, 1, 4, 9, 5, 2, 6, 7, 3, 8}
	if !sliceEquals(expected[:], list[:]) {
		t.Errorf("Array is incorrect, got: %v, want: %v.", list, expected)
	}
}

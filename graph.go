package graph

import (
	"github.com/mcustiel/queue"
)

type Graph struct {
	root *Node
}

func New(rootNode *Node) *Graph {
	graph := new(Graph)
	graph.root = rootNode
	return graph
}

func (graph *Graph) Walk(callback func(*Node)) {
	visited := make(map[*Node]bool)
	var function func(node *Node)

	function = func(node *Node) {
		callback(node)
		visited[node] = true
		for e := node.adyacent.Front(); e != nil; e = e.Next() {
			if wasVisited := visited[e.Value.(*Node)]; !wasVisited {
				function(e.Value.(*Node))
			}
		}
	}
	function(graph.root)
}

func (graph *Graph) Bfs(callback func(*Node)) {
	visited := make(map[*Node]bool)
	queue := queue.NewQueue()
	queue.Add(graph.root)
	visited[graph.root] = true

	for !queue.Empty() {
		current := queue.Get().(*Node)
		iterator := current.GetIterator()
		for currentNode := iterator(); currentNode != nil; currentNode = iterator() {
			if _, wasVisited := visited[currentNode]; !wasVisited {
				visited[currentNode] = true
				queue.Add(currentNode)
				callback(currentNode)
			}
		}
	}
}

func (graph *Graph) Dfs(callback func(*Node)) {
	visited := make(map[*Node]bool)
	var function func(*Node)

	function = func(node *Node) {
		if _, wasVisited := visited[node]; !wasVisited {
			visited[node] = true
			callback(node)
			iterator := node.GetIterator()
			for currentNode := iterator(); currentNode != nil; currentNode = iterator() {
				function(currentNode)
			}
		}
	}
	function(graph.root)
}

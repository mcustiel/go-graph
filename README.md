# go-graph
Implements a general graph data structure in Go

This project is part of my "learn go" projects

##  Types

* graph.Node
* graph.Graph

## Functions

* `graph.New(*Node) *Graph` - Creates a new graph using the given node as start point for traversing methods.
* `graph.NewNode(interface{}) *Node` - Creates a new node contained a given value.

## Methods

### Node

* `(*Node) AddAdyacent(*Node)` - Makes the given node adyacent to the current one
* `(*Node) RemoveAdyacent(*Node)` - Removes the given node from the current one. (TODO: return the removed node)
* `(*Node) GetIterator() func() *Node` - Returns an iterator function that will return the next adyacent node each time is called until nil.
* `(*Node) Value() interface{}` - Returns the value stored in the current node. You should convert it to the expected type.
* `(*Node) Adyacents() *list.List` - Returns the adyacents node as a Go's double linked list.
 
### Graph

* `(*Graph) Walk(func(*Node))` - Walks all the graph nodes without specific order and executes the given function for each of them.
* `(*Graph) Bfs(func(*Node))` - Traverses the graph using BFS algorithm and executes the given function for each node.
* `(*Graph) Dfs(func(*Node))` - Traverses the graph using DFS algorithm and executes the given function for each node.

## Example

### Saves a directories tree into a graph:

```go
// GetDirectoryContentents(dirName) should be implemented returning []os.FileInfo

func ScanDirToTree(dirName string) *graph.Graph {
	var function func(string) *graph.Node
	function = func(curDir string) *graph.Node {
		node := graph.NewNode(curDir)
		for _, file := range GetDirectoryContents(curDir) {
			if file.IsDir() {
				node.AddAdyacent(function(curDir + string(os.PathSeparator) + file.Name()))
			}
		}
		return node
	}
	return graph.New(function(dirName))
}
```

### Displaying the folders in the graph using BFS:

```go
graphInstance := ScanDirToTree("some/directory/path")
graphInstance.Bfs(func(node *graph.Node) {
	fmt.Println(node.Value().(string))
})
```

 

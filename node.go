package graph

import (
	"container/list"
)

type Node struct {
	value    interface{}
	adyacent *list.List
}

func NewNode(nodeValue interface{}) *Node {
	node := new(Node)
	node.value = nodeValue
	node.adyacent = list.New()
	return node
}

func (node *Node) AddAdyacent(adyacentNode *Node) {
	node.adyacent.PushBack(adyacentNode)
}

func (node *Node) RemoveAdyacent(adyacent *Node) {
	for e := node.adyacent.Front(); e != nil; e = e.Next() {
		current := e.Value.(*Node)
		if current == adyacent {
			tmp := e.Next()
			node.adyacent.Remove(e)
			e = tmp
		}
	}
}

func (node *Node) GetIterator() func() *Node {
	current := node.adyacent.Front()
	return func() *Node {
		if current != nil {
			value := current.Value.(*Node)
			current = current.Next()
			return value
		}
		return nil
	}
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Adyacents() *list.List {
	return node.adyacent
}

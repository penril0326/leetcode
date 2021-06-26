package graph

import (
	"container/list"
)

func NewGraph() *Graph {
	g := new(Graph)
	g.nodeList = make(map[int](*list.List))

	return g
}

func (g *Graph) AddNode(index, value int) *Node {
	n := new(Node)
	n.index = index
	n.value = value
	n.visit = false

	l := new(list.List)
	g.nodeList[index] = l

	return n
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	g.nodeList[n1.index].PushBack(n2)
	g.nodeList[n2.index].PushBack(n1)
}

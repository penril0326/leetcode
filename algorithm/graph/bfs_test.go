package graph

import (
	"testing"
)

func TestGraph_bfs(t *testing.T) {
	g := NewGraph()

	data := []int{100, 200, 100, 500, 100, 600}
	nodeList := []*Node{}
	for idx, v := range data {
		temp := g.AddNode(idx+1, v)
		nodeList = append(nodeList, temp)
	}

	g.AddEdge(nodeList[0], nodeList[1])
	g.AddEdge(nodeList[1], nodeList[2])
	g.AddEdge(nodeList[1], nodeList[4])
	g.AddEdge(nodeList[3], nodeList[4])
	g.AddEdge(nodeList[4], nodeList[5])

	g.bfs(nodeList[0])
}

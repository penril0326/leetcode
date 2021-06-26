package graph

import (
	"container/list"
	"fmt"
)

func (g *Graph) bfs(start *Node) {
	q := list.New()
	q.PushBack(start)
	start.visit = true

	fmt.Printf("%v -> ", start.index)

	for n := q.Front(); n != nil; n = n.Next() {
		node := n.Value.(*Node)
		for e := g.nodeList[node.index].Front(); e != nil; e = e.Next() {
			neighbor := e.Value.(*Node)
			if neighbor.visit == false {
				q.PushBack(neighbor)
				neighbor.visit = true
				fmt.Printf("%v -> ", neighbor.index)
			}
		}
	}

	fmt.Printf("nil")
}

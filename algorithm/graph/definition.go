package graph

import "container/list"

type Graph struct {
	nodeList map[int](*list.List)
}

type Node struct {
	index int
	value int
	visit bool
}

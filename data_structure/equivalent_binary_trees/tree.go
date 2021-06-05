package equivalent_tree

import (
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	recursiveWalk(t, ch)
	close(ch)
}

func recursiveWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		recursiveWalk(t.Left, ch)
		ch <- t.Value
		recursiveWalk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		n1, ok1 := <-c1
		n2, ok2 := <-c2

		if (ok1 != ok2) || (n1 != n2) {
			return false
		}

		if (ok1 == false) && (ok2 == false) {
			break
		}
	}

	return true
}

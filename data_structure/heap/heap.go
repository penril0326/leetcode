package heap

type MyHeap interface {
	Push(v interface{})
	Pop() interface{}
	Len() int
	Top() interface{}
}

type mHeap struct {
	h   []interface{}
	cmp func(i, j interface{}) bool
}

func NewHeap(vals []interface{}, cmpFunc func(i, j interface{}) bool) MyHeap {
	h := mHeap{
		h:   vals,
		cmp: cmpFunc,
	}

	if h.cmp == nil {
		return nil
	}

	if h.Len() > 0 {
		h.heapify()
	}

	return &h
}

func (h mHeap) Len() int {
	return len(h.h)
}

func (h mHeap) Top() interface{} {
	if h.Len() > 0 {
		return h.h[0]
	}

	return nil
}

func (h *mHeap) Push(v interface{}) {
	h.h = append(h.h, v)
	h.bubbleUp(h.Len() - 1)
}

func (h *mHeap) bubbleUp(idx int) {
	parentIdx := (idx - 1) / 2
	for (idx > 0) && h.cmp(h.h[idx], h.h[parentIdx]) {
		h.swap(idx, parentIdx)
		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}
}

func (h *mHeap) swap(i, j int) {
	h.h[i], h.h[j] = h.h[j], h.h[i]
}

func (h *mHeap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}

	root := h.h[0]
	lastIdx := h.Len() - 1
	h.swap(0, lastIdx)
	h.h = h.h[:lastIdx]
	h.bubbleDown(0)

	return root
}

func (h *mHeap) bubbleDown(idx int) {
	lcIdx := 2*idx + 1
	rcIdx := 2*idx + 2
	cur := idx

	if (lcIdx < h.Len()) && h.cmp(h.h[lcIdx], h.h[cur]) {
		cur = lcIdx
	}

	if (rcIdx < h.Len()) && h.cmp(h.h[rcIdx], h.h[cur]) {
		cur = rcIdx
	}

	if cur != idx {
		h.swap(cur, idx)
		h.bubbleDown(cur)
	}
}

func (h *mHeap) heapify() {
	lastIdx := h.Len() - 1
	for start := (lastIdx - 1) / 2; start >= 0; start-- {
		h.bubbleDown(start)
	}
}

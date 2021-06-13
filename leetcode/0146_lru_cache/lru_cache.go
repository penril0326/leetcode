package lru_cache

type LRUCache struct {
	capacity int
	hash     map[int]*linkList
	head     *linkList
	tail     *linkList
}

type linkList struct {
	key      int
	value    int
	previous *linkList
	next     *linkList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*linkList),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, isExist := this.hash[key]; isExist == true {
		this.removeNode(node)
		this.insert(node)

		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, isExist := this.hash[key]; isExist == true {
		node.value = value
		this.removeNode(node)
		this.insert(node)
	} else {
		newNode := &linkList{
			key:      key,
			value:    value,
			next:     nil,
			previous: nil,
		}

		this.insert(newNode)

		if len(this.hash) > this.capacity {
			delete(this.hash, this.tail.key)
			this.removeNode(this.tail)
		}
	}
}

func (this *LRUCache) insert(node *linkList) {
	node.previous = nil
	node.next = this.head

	if this.head != nil {
		this.head.previous = node
	}

	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}

	this.head = node
	this.hash[node.key] = node
}

func (this *LRUCache) removeNode(node *linkList) {
	if node == this.head {
		this.head = node.next
		node.next = nil

		return
	}

	if node == this.tail {
		this.tail = node.previous
		this.tail.next = nil
		node.previous = nil

		return
	}

	node.previous.next = node.next
	node.next.previous = node.previous
}

package main

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next, tail.Pre = tail, head

	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.moveToHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Value = value
		this.moveToHead(node)
		return
	}

	if len(this.Cache) == this.Capacity {
		lastKey := this.deleteLast()
		delete(this.Cache, lastKey)
	}

	newNode := &Node{
		Key:   key,
		Value: value,
	}
	this.addToHead(newNode)
	this.Cache[key] = newNode
}

func (this *LRUCache) moveToHead(node *Node) {
	node.Next.Pre, node.Pre.Next = node.Pre, node.Next
	node.Next = this.Head.Next
	node.Next.Pre = node
	this.Head.Next = node
	node.Pre = this.Head
}

func (this *LRUCache) deleteLast() int {
	last := this.Tail.Pre
	this.Tail.Pre.Pre.Next = this.Tail
	this.Tail.Pre = this.Tail.Pre.Pre
	return last.Key
}

func (this *LRUCache) addToHead(node *Node) {
	node.Next = this.Head.Next
	node.Pre = this.Head
	this.Head.Next.Pre = node
	this.Head.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

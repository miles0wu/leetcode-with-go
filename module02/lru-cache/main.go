package main

type LRUCache struct {
	head     *Node
	tail     *Node
	address  map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := new(Node)
	tail := new(Node)
	head.Next, head.Prev = tail, nil
	tail.Next, tail.Prev = nil, head
	return LRUCache{head, tail, map[int]*Node{}, capacity}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.address[key]
	if !ok {
		return -1
	}

	if this.head.Next != node {
		// move node to head
		remove(node)
		insert(this.head, node)
	}

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.address[key]; ok {
		node.Value = value
		if this.head.Next != node {
			remove(node)
			insert(this.head, node)
		}
		return
	}

	// new key
	if len(this.address) >= this.capacity {
		k := this.tail.Prev.Key
		remove(this.tail.Prev)
		delete(this.address, k)
	}

	node := NewNode(key, value)
	insert(this.head, node)
	this.address[key] = node
}

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

func NewNode(key, value int) *Node {
	return &Node{key, value, nil, nil}
}

func insert(before, node *Node) {
	node.Next, before.Next.Prev = before.Next, node
	before.Next, node.Prev = node, before
}

func remove(node *Node) {
	node.Next.Prev, node.Prev.Next = node.Prev, node.Next
	node.Prev, node.Next = nil, nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	seq := [][]int{}
	if root == nil {
		return seq
	}

	// create queue and enq root
	queue := []*NodeAndLayer{NewNodeAndLayer(root, 0)}
	for len(queue) != 0 {
		// deq
		node := queue[0]
		queue = queue[1:]
		if node.Depth > len(seq)-1 {
			seq = append(seq, []int{})
		}
		seq[node.Depth] = append(seq[node.Depth], node.Node.Val)

		// enq children
		for _, child := range node.Node.Children {
			queue = append(queue, NewNodeAndLayer(child, node.Depth+1))
		}
	}

	return seq
}

type NodeAndLayer struct {
	Node  *Node
	Depth int
}

func NewNodeAndLayer(node *Node, layer int) *NodeAndLayer {
	return &NodeAndLayer{node, layer}
}

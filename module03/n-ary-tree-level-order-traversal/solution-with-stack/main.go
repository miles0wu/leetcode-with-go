package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	seq := []int{}

	if root == nil {
		return seq
	}

	// push root
	stack := []*Node{root}

	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		seq = append(seq, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			// push childrens
			stack = append(stack, node.Children[i])
		}
	}

	return seq
}

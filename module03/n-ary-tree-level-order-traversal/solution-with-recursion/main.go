package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	seq := []int{}
	var dfs func(*Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		seq = append(seq, root.Val)
		for _, child := range root.Children {
			dfs(child)
		}
	}
	dfs(root)

	return seq
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(int, int, int, int) *TreeNode
	// 前序: preorder[l1:r1]
	// 中序: preorder[l2:r2]
	build = func(l1, r1, l2, r2 int) *TreeNode {
		if l1 > r1 {
			return nil
		}
		root := &TreeNode{preorder[l1], nil, nil}
		// 求root在中序的位置mid，inorder[mid]等於root
		mid := l2
		for inorder[mid] != root.Val {
			mid++
		}

		// 前序左子樹range: l1 + 1 ~ l1 + (mid-l2-1) + 1
		// 中序左子樹range: l2 ~ mid-1
		root.Left = build(l1+1, l1+mid-l2, l2, mid-1)

		// 前序右子樹range: l1 + (mid-l2) + 1 ~ r2
		// 中序右子樹range: mid + 1 ~ r2
		root.Right = build(l1+mid-l2+1, r1, mid+1, r2)
		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

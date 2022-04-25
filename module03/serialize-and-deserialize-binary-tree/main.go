package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	seq := []string{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			seq = append(seq, "nil")
			return
		}
		val := strconv.Itoa(root.Val)
		seq = append(seq, val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(seq, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	curr := 0
	values := strings.Split(data, ",")

	var recur func() *TreeNode
	recur = func() *TreeNode {
		value := values[curr]
		if value == "nil" {
			curr++
			return nil
		}
		intval, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		curr++

		root := new(TreeNode)
		root.Val = intval
		root.Left = recur()
		root.Right = recur()

		return root
	}
	return recur()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

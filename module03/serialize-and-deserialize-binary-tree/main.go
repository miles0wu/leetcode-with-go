package main

import (
	"fmt"
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
	curr := 1
	if root != nil {
		seq = append(seq, strconv.Itoa(root.Val))
	}
	fmt.Println(seq)

	var recur func(*TreeNode, int)
	recur = func(root *TreeNode, depth int) {
		// 直到depth這層，應該有多少個元素，第0層1個，第1層3個...
		totals := (1 << (depth + 1)) - 1
		if len(seq) != totals {
			for i := len(seq); i < totals; i++ {
				seq = append(seq, "nil")
			}
		}
		if root == nil {
			curr++
			return
		}

		if root.Left != nil {
			seq[(curr)*2-1] = strconv.Itoa(root.Left.Val)
		}
		if root.Right != nil {
			seq[(curr)*2] = strconv.Itoa(root.Right.Val)
		}
		curr++

		recur(root.Left, depth+1)
		recur(root.Right, depth+1)
	}
	recur(root, 1)
	fmt.Println(seq)
	return strings.Join(seq, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	fmt.Println(data)
	values := strings.Split(data, ",")
	if values[0] == "nil" {
		return nil
	}
	val, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}
	root := &TreeNode{val, nil, nil}

	treeMap := map[int]*TreeNode{0: root}
	// 只需遍歷到totals/2
	for i := 1; i <= len(values)/2; i++ {
		parent := treeMap[(i+1)/2-1]
		if parent == nil {
			continue
		}

		var node *TreeNode
		if values[i] != "nil" {
			val, err := strconv.Atoi(values[i])
			if err != nil {
				panic(err)
			}
			node = &TreeNode{val, nil, nil}
		}

		// 左子樹正好可被2整除
		if (i+1)%2 == 0 {
			parent.Left = node
		} else {
			parent.Right = node
		}
		treeMap[i] = node
	}

	for i, tree := range treeMap {
		if tree == nil {
			fmt.Printf("i=%d, v=nil\n", i)
		} else {
			fmt.Printf("i=%d, v=%v\n", i, tree.Val)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

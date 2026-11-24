package main

import (
	"fmt"
	"math"
)

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// IsValidBST reports whether the tree satisfies the strict BST property:
// every node is greater than everything in its left subtree and smaller
// than everything in its right. Each node is checked against a (lo, hi)
// window narrowed by its ancestors.
func IsValidBST(root *TreeNode) bool {
	var valid func(n *TreeNode, lo, hi int) bool
	valid = func(n *TreeNode, lo, hi int) bool {
		if n == nil {
			return true
		}
		if n.Val <= lo || n.Val >= hi {
			return false
		}
		return valid(n.Left, lo, n.Val) && valid(n.Right, n.Val, hi)
	}
	return valid(root, math.MinInt, math.MaxInt)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	ok := tn(2, tn(1, nil, nil), tn(3, nil, nil))
	bad := tn(5, tn(1, nil, nil), tn(4, tn(3, nil, nil), tn(6, nil, nil)))
	fmt.Println(IsValidBST(ok), IsValidBST(bad))
}

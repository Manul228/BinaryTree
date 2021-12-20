package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pt(root *TreeNode, travers *[]int) {
	if root == nil {
		return
	}

	*travers = append(*travers, root.Val)

	if root.Left != nil {
		pt(root.Left, travers)
	}
	if root.Right != nil {
		pt(root.Right, travers)
	}
}

func preorderTraversal(root *TreeNode) []int {
	var travers []int

	pt(root, &travers)

	return travers
}

func it(root *TreeNode, travers *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		it(root.Left, travers)
	}

	*travers = append(*travers, root.Val)

	if root.Right != nil {
		it(root.Right, travers)
	}
}

func inorderTraversal(root *TreeNode) []int {
	var travers []int

	it(root, &travers)

	return travers
}

func postt(root *TreeNode, travers *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		postt(root.Left, travers)
	}

	if root.Right != nil {
		postt(root.Right, travers)
	}

	*travers = append(*travers, root.Val)

}

func postorderTraversal(root *TreeNode) []int {
	var travers []int

	postt(root, &travers)

	return travers
}

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int

	if root == nil {
		return levels
	}

	var queue []*TreeNode
	var leaves []*TreeNode
	var node *TreeNode
	var level []int

	queue = append(queue, root)

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		level = append(level, node.Val)

		if node.Left != nil {
			leaves = append(leaves, node.Left)
		}

		if node.Right != nil {
			leaves = append(leaves, node.Right)
		}

		if len(queue) == 0 {
			levels = append(levels, level)
			level = make([]int, 0)

			for _, l := range leaves {
				queue = append(queue, l)
			}

			leaves = make([]*TreeNode, 0)
		}
	}

	return levels
}

func md(root *TreeNode, depth int, answer *int) {
	if root == nil {
		if depth > *answer {
			*answer = depth
		}
		return
	}

	md(root.Left, depth+1, answer)
	md(root.Right, depth+1, answer)

	return
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var answer int

	md(root, 0, &answer)

	return answer
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil {
		if left.Val == right.Val {
			return isMirror(left.Right, right.Left) && isMirror(left.Left, right.Right)
		}
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
		return false
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func search(arr []int, start, end int, value int) int {
	var i int
	for i = start; i < end; i++ {
		if arr[i] == value {
			break
		}
	}
	return i
}

func buildUtil(in []int, post []int, inStart, inEnd int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	var node *TreeNode = &TreeNode{}
	node.Val = post[postEnd]

	if inStart == inEnd {
		return node
	}

	inIndex := search(in, inStart, inEnd, node.Val)

	node.Left = buildUtil(in, post, inStart, inIndex-1, postStart, postStart-inStart+inIndex-1)
	node.Right = buildUtil(in, post, inIndex+1, inEnd, postEnd-inEnd+inIndex, postEnd-1)

	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	tree := buildUtil(inorder, postorder, 0, n-1, 0, n-1)
	return tree
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	left1 := TreeNode{}
	left1.Val = 4

	right1 := TreeNode{}
	right1.Val = 5

	// left2 := TreeNode{}
	// left2.Val = 4

	right2 := TreeNode{}
	right2.Val = 8

	root.Left = &left
	root.Right = &right

	root.Left.Left = &left1
	root.Left.Right = &right1

	// root.Right.Left = &left2
	root.Right.Right = &right2

	in := inorderTraversal(&root)
	post := postorderTraversal(&root)

	result := buildTree(in, post)

	fmt.Println(result)
}

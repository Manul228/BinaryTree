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

func hps(root *TreeNode, currentSum int, targetSum *int, isFinded *bool) {

	if *isFinded {
		return
	}

	var newCurrentSum int = currentSum + root.Val

	if newCurrentSum == *targetSum && root.Left == nil && root.Right == nil {
		*isFinded = true
	}

	if root.Left != nil {
		hps(root.Left, newCurrentSum, targetSum, isFinded)
	}

	if root.Right != nil {
		hps(root.Right, newCurrentSum, targetSum, isFinded)
	}

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var currentSum int = 0
	var isFinded bool = false
	hps(root, currentSum, &targetSum, &isFinded)
	return isFinded
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

	result := hasPathSum(&root, 11)

	fmt.Println(result)
}

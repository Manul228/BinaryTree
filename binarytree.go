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

func helperMaxDepth(root *TreeNode, depth int, answer *int) {
	if root == nil {
		if depth > *answer {
			*answer = depth
		}
		return
	}

	helperMaxDepth(root.Left, depth+1, answer)
	helperMaxDepth(root.Right, depth+1, answer)

	return
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var answer int

	helperMaxDepth(root, 0, &answer)

	return answer
}

func levelOrderNodes(root *TreeNode) [][]*TreeNode {
	var levels [][]*TreeNode

	if root == nil {
		return levels
	}

	var queue []*TreeNode
	var leaves []*TreeNode
	var node *TreeNode
	var level []*TreeNode

	queue = append(queue, root)

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		level = append(level, node)

		if node.Left != nil {
			leaves = append(leaves, node.Left)
		}

		if node.Right != nil {
			leaves = append(leaves, node.Right)
		}

		if len(queue) == 0 {
			levels = append(levels, level)
			level = make([]*TreeNode, 0)

			for _, l := range leaves {
				queue = append(queue, l)
			}

			leaves = make([]*TreeNode, 0)
		}
	}

	return levels
}

func isNodesEqual(left, right *TreeNode) bool {
	if left.Val != right.Val {

	}
}

func isSymmetric(root *TreeNode) bool {
	var leftTraverse [][]*TreeNode = levelOrderNodes(root.Left)
	var rightTraverse [][]*TreeNode = levelOrderNodes(root.Right)

	if len(leftTraverse) != len(rightTraverse) {
		return false
	}

	fmt.Println(leftTraverse)
	fmt.Println(rightTraverse)

	for i, lt := range leftTraverse {
		for j, rt := range rightTraverse[i] {
			if lt[len(lt)-j-1] != rt {
				return false
			}
		}
	}

	return true
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	// left1 := TreeNode{}
	// left1.Val = 3

	right1 := TreeNode{}
	right1.Val = 4

	left2 := TreeNode{}
	left2.Val = 4

	// right2 := TreeNode{}
	// right2.Val = 3

	root.Left = &left
	root.Right = &right

	// root.Left.Left = &left1
	root.Left.Right = &right1

	root.Right.Left = &left2
	// root.Right.Right = &right2

	result := isSymmetric(&root)

	fmt.Println(result)
}

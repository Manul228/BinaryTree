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

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 3

	root.Left = &left
	root.Right = &right

	result := postorderTraversal(&root)

	fmt.Println(result)
}

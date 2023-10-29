package pattern

import (
	"fmt"
)

type BreadthFirstSearch struct{}

func InitBreadthFirstSearch() BreadthFirstSearch {
	return BreadthFirstSearch{}
}

/**
 * Binary Tree Level Order Traversal
 *
 * Input: [3,9,20,null,null,15,7]
 * Output: [[3],[9,20],[15,7]]
 */
func (p *BreadthFirstSearch) BinaryTreeLevelOrderTraversal(root *TreeNode) {
	results := [][]int{}

	var bfs func(*TreeNode, int)
	bfs = func(node *TreeNode, i int) {
		if node == nil {
			return
		}

		if len(results)-1 < i {
			results = append(results, []int{})
		}

		results[i] = append(results[i], node.Val)

		if node.Left != nil {
			bfs(node.Left, i+1)
		}
		if node.Right != nil {
			bfs(node.Right, i+1)
		}
	}
	bfs(root, 0)

	fmt.Println("breadth first search: binary tree level order traversal")
	fmt.Println("input:", root)
	fmt.Println("output:", results)
}

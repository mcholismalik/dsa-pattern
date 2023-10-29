package pattern

import (
	"fmt"
	"math"
)

type DepthFirstSearch struct{}

func InitDepthFirstSearch() DepthFirstSearch {
	return DepthFirstSearch{}
}

/**
 * Binary Tree Maximum Path Sum
 *
 * Input: [1,2,3]
 * Output: 6 (2 + 1 + 1)
 *
 * Input: [-10,9,20,null,null,15,7]
 * Output: 42 (15 + 20 + 7)
 *
 * Input: [1,2,3,null,null,4,5]
 * Output: 12
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (p *DepthFirstSearch) BinaryTreeMaximumPathSum(root *TreeNode) {
	result := math.MinInt

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMaxSum := max(dfs(node.Left), 0)
		rightMaxSum := max(dfs(node.Right), 0)
		treeSum := node.Val + leftMaxSum + rightMaxSum
		result = max(treeSum, result)

		return node.Val + max(leftMaxSum, rightMaxSum)
	}
	dfs(root)

	fmt.Println("depth first search: binary tree maximum path sum")
	fmt.Println("input:", root)
	fmt.Println("output:", result)
}

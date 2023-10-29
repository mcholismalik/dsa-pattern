package main

import "pg-go/pattern"

func main() {
	// sliding window
	// substring
	// backtracking
	// greedy
	// dynamic programming
	// two pointers
	// binary search
	// graph

	// array string
	// as := pattern.InitArrayString()
	// as.InsertDeleteGetRandomO1()
	// as.RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	// as.RotateArray([]int{-1, -100, 3, 99}, 2)
	// as.RotateArray([]int{1, 2}, 3)
	// as.RotateArray([]int{1, 2}, 5)
	// as.RotateArray([]int{1, 2, 3}, 4)

	// two pointers
	// tp := pattern.InitTwoPointers()
	// tp.TwoSum2([]int{2, 7, 11, 15}, 9)
	// tp.TwoSum2([]int{2, 3, 4}, 6)

	// sliding window
	// sw := pattern.InitSlidingWindow()
	// sw.LengthOfLongestSubstring("abcabcbb")
	// sw.LengthOfLongestSubstring("abba")

	// greedy
	// sw := pattern.InitGreedy()
	// sw.JumpGame([]int{2, 3, 1, 1, 4})
	// sw.JumpGame([]int{3, 2, 1, 0, 4})

	// backtracking
	bt := pattern.InitBacktracking()
	// bt.Pemutations([]int{1, 2, 3})
	bt.CombinationSum([]int{2, 3, 6, 7}, 7)
	// bt.CombinationSum2([]int{2, 3, 6, 7}, 7)
	// bt.CombinationSum([]int{2, 3, 5}, 8)

	// dynamic programming
	// dp := pattern.InitDynamicProgramming()
	// dp.TriangleMinTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	// dp.TriangleMinTotal([][]int{{-10}})

	// hashmap
	// hm := pattern.InitHashmap()
	// hm.GroupAnagram([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	// hm.GroupAnagram([]string{""})

	// depth first search
	// dfs := pattern.InitDepthFirstSearch()
	// dfs.BinaryTreeMaximumPathSum(&pattern.TreeNode{
	// 	Val: -10,
	// 	Left: &pattern.TreeNode{
	// 		Val: 9,
	// 	},
	// 	Right: &pattern.TreeNode{
	// 		Val: 20,
	// 		Left: &pattern.TreeNode{
	// 			Val: 15,
	// 		},
	// 		Right: &pattern.TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// })
	// dfs.BinaryTreeMaximumPathSum(&pattern.TreeNode{
	// 	Val: -3,
	// })

	// breadth first search
	// bfs := pattern.InitBreadthFirstSearch()
	// bfs.BinaryTreeLevelOrderTraversal(&pattern.TreeNode{
	// 	Val: 3,
	// 	Left: &pattern.TreeNode{
	// 		Val: 9,
	// 	},
	// 	Right: &pattern.TreeNode{
	// 		Val: 20,
	// 		Left: &pattern.TreeNode{
	// 			Val: 15,
	// 		},
	// 		Right: &pattern.TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// })
}

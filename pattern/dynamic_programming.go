package pattern

import (
	"fmt"
)

type DynamicProgramming struct{}

func InitDynamicProgramming() DynamicProgramming {
	return DynamicProgramming{}
}

/**
 * Triangle Minimum Total
 *
 * Input: [[2],[3,4],[6,5,7],[4,1,8,3]]
 * Output: 11
 *
 * Input: [[-10]]
 * Output: -10
 */
func (p *DynamicProgramming) TriangleMinTotal(nums [][]int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(nums)
	dp := make([]int, n+1)
	for k := n - 1; k >= 0; k-- {
		for i := 0; i <= k; i++ {
			dp[i] = nums[k][i] + min(dp[i], dp[i+1])
		}
	}
	result := dp[0]

	fmt.Println("dynamic programming: triangle minimum total")
	fmt.Println("input:", nums)
	fmt.Println("output:", result)
}

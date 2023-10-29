package pattern

import (
	"fmt"
)

type Greedy struct{}

func InitGreedy() Greedy {
	return Greedy{}
}

/**
 * Jump Game
 *
 * Input: [2,3,1,1,4]
 * Output: true
 *
 * Input: [3,2,1,0,4]
 * Output: false
 */
func (p *Greedy) JumpGame(nums []int) {
	result := false

	goal := len(nums) - 1
	i := goal
	for i >= 0 {
		if i+nums[i] >= goal {
			goal = i
		}
		i--
	}

	result = goal == 0

	fmt.Println("greedy: jump game")
	fmt.Println("input:", nums)
	fmt.Println("output:", result)
}

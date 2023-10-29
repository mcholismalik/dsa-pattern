package pattern

import (
	"fmt"
)

type Backtracking struct{}

func InitBacktracking() Backtracking {
	return Backtracking{}
}

/**
 * Permutations
 *
 * Input: [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 * Input: [0,1]
 * Output: [[0,1],[1,0]]
 */
func (p *Backtracking) Pemutations(nums []int) {
	results := make([][]int, 0)

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			results = append(results, perm)
			return
		}

		for i := index; i < len(nums); i++ {
			nums[index], nums[i] = nums[i], nums[index]

			backtrack(index + 1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	backtrack(0)

	fmt.Println("backtracking: permutations")
	fmt.Println("input:", nums)
	fmt.Println("output:", results)
}

/**
 * Combination Sum
 *
 * Input: nums = [2,3,6,7], target = 7
 * Output: [[2,2,3],[7]]
 *
 * Input: nums = [2,3,5], target = 8
 * Output: [[2,2,2,2],[2,3,3],[3,5]]
 */
func (p *Backtracking) CombinationSum(nums []int, target int) {
	results := [][]int{}

	var bt func(i, sum int, comb []int)
	bt = func(i, sum int, comb []int) {
		if i == len(nums) {
			return
		}

		sum += nums[i]
		newComb := append([]int{}, comb...)
		newComb = append(newComb, nums[i])

		if sum == target {
			results = append(results, newComb)
		}
		if sum < target {
			bt(i, sum, newComb)
		}
		bt(i+1, sum-nums[i], comb)
	}

	bt(0, 0, []int{})

	fmt.Println("backtracking: combination sum 2")
	fmt.Println("input:", nums)
	fmt.Println("output:", results)
}

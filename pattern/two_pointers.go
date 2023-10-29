package pattern

import (
	"fmt"
)

type TwoPointers struct{}

func InitTwoPointers() TwoPointers {
	return TwoPointers{}
}

/**
 * Two Sum 2 - Input Array Is Sorted
 *
 * Input: [2,7,11,15], 9
 * Output: [1,2]
 *
 * Input: [2,3,4], 6
 * Output: [1,3]
 */
func (p *TwoPointers) TwoSum2(nums []int, target int) {
	results := []int{}

	l, r := 0, 1
	for len(results) < 2 {
		sum := nums[l] + nums[r]
		if sum == target {
			results = []int{l + 1, r + 1}
		}

		if r == len(nums)-1 {
			l++
			r = l
		}
		if l == len(nums)-1 {
			break
		}
		r++
	}

	fmt.Println("two pointer: two sum 2 - input array is sorted")
	fmt.Println("input:", nums, target)
	fmt.Println("output:", results)
}

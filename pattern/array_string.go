package pattern

import (
	"fmt"
	"math/rand"
)

type ArrayString struct{}

func InitArrayString() ArrayString {
	return ArrayString{}
}

/**
 * RandomizedSet
 *
 * Input:
 * ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
 * [[], [1], [2], [2], [], [1], [2], []]
 * Output: [null, true, false, true, 2, true, false, 2]
 */

type RandomizedSet struct {
	mapValues map[int]int
	indexes   []int
}

func initRandomizedSet() RandomizedSet {
	return RandomizedSet{
		mapValues: make(map[int]int),
		indexes:   make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.mapValues[val]
	if ok {
		return false
	}

	rs.indexes = append(rs.indexes, val)
	rs.mapValues[val] = len(rs.indexes) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	mapVal, ok := rs.mapValues[val]
	if !ok {
		return false
	}

	lastVal := rs.indexes[len(rs.indexes)-1]
	rs.mapValues[lastVal] = mapVal
	rs.indexes[mapVal] = lastVal

	delete(rs.mapValues, val)
	rs.indexes = rs.indexes[:len(rs.indexes)-1]

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randVal := rand.Intn(len(rs.indexes)) + 0
	return rs.indexes[randVal]
}

func (p *ArrayString) InsertDeleteGetRandomO1() {
	rs := initRandomizedSet()

	outputs := []interface{}{
		rs.Insert(1),
		rs.Remove(2),
		rs.Insert(2),
		rs.GetRandom(),
		rs.Remove(1),
		rs.Insert(2),
		rs.GetRandom(),
	}

	fmt.Println("array string: insert delete get random o(1)")
	fmt.Println("input:", 12212)
	fmt.Println("output:", outputs)
}

/**
 * Rotate Array
 *
 * Input: [1,2,3,4,5,6,7], 3
 * Output: [5,6,7,1,2,3,4]
 */
func (p *ArrayString) RotateArray(nums []int, k int) {
	tempNums := make([]int, len(nums))
	copy(tempNums, nums)

	mapNums := make(map[int]int)
	for i := range nums {
		mapNums[i] = nums[i]
	}

	maxLen := len(nums)
	maxI := maxLen - 1
	ptr := maxLen - k
	for ptr < 0 {
		ptr += maxLen
	}

	for i := range nums {
		nums[i] = mapNums[ptr]
		ptr += 1
		if ptr > maxI {
			ptr -= maxLen
		}
	}

	fmt.Println("array string: rotate array")
	fmt.Println("input:", tempNums, k)
	fmt.Println("output:", nums)
}

/**
 * Rotate Array v2
 *
 * Input: [1,2,3,4,5,6,7], 3
 * Output: [5,6,7,1,2,3,4]
 */
func (p *ArrayString) RotateArrayV2(nums []int, k int) {
	tempNums := make([]int, len(nums))
	copy(tempNums, nums)

	n := len(nums)
	k = k % n // Ensure k is within bounds

	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

	fmt.Println("array string: rotate array")
	fmt.Println("input:", tempNums, k)
	fmt.Println("output:", nums)
}

package pattern

import (
	"fmt"
)

type Hashmap struct{}

func InitHashmap() Hashmap {
	return Hashmap{}
}

/**
 * Group Anagram
 *
 * Input: ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * Input: [""]
 * Output: [[""]]
 */
func (p *Hashmap) GroupAnagram(strs []string) {
	results := [][]string{}

	type CharCount [26]int
	hash := make(map[CharCount][]string)
	for i := range strs {
		charCount := CharCount{}
		for j := range strs[i] {
			charCount[strs[i][j]-'a']++
		}

		hash[charCount] = append(hash[charCount], strs[i])
	}

	for i := range hash {
		results = append(results, hash[i])
	}

	fmt.Println("backtracking: group anagram")
	fmt.Println("input:", strs)
	fmt.Println("output:", results)
}

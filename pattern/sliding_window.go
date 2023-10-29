package pattern

import "fmt"

type SlidingWindow struct{}

func InitSlidingWindow() SlidingWindow {
	return SlidingWindow{}
}

/**
 * Length Of Longest Substring
 *
 * Input: abcabcbb
 * Output: 3 (abc)
 *
 * Input: abba
 * Output: 2 (ab)
 */
func (p *SlidingWindow) LengthOfLongestSubstring(s string) {
	left := 0
	max := 0
	mapCharIndex := make(map[string]int)

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		val, ok := mapCharIndex[char]
		if ok && val >= left {
			left = val + 1
		}

		count := i - left + 1
		if count >= max {
			max = count
		}

		mapCharIndex[char] = i
	}

	fmt.Println("sliding window: length of longest substring")
	fmt.Println("input:", s)
	fmt.Println("output:", max)
}

package main

import "math"

func main() {

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := math.MinInt64
	win := map[byte]int{}
	left, right := 0, 0
	for right < len(s) {
		r := s[right]
		right++
		win[r]++
		for win[r] > 1 {
			l := s[left]
			left++
			win[l]--
		}
		ans = max(right-left, ans)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

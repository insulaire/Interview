package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getMaxLenth("asdccdddaaddeccc", 2))
}

//给定一个字符串，请你找出其中重复K次字符的 最长子串 的长度。
// s=aaabb  k=2   return 4 aabb
func getMaxLenth(s string, k int) int {
	return helper(0, len(s)-1, k, s)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func helper(start, end, k int, s string) int {
	if end-start+1 < k {
		return 0
	}
	cut := map[byte]int{}
	for i := start; i <= end; i++ {
		cut[s[i]]++
	}
	for end-start+1 >= k && cut[s[start]] != k {
		cut[s[start]]--
		start++
	}
	for end-start+1 >= k && cut[s[end]] != k {
		cut[s[end]]--
		end--
	}

	if end-start+1 < k {
		return 0
	}

	for i := start; i <= end; i++ {
		if cut[s[i]] != k {
			return max(helper(start, i-1, k, s), helper(i+1, end, k, s))
		}
	}
	return end - start + 1
}

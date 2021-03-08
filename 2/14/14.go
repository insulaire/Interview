package main

import "fmt"

func main() {
	fmt.Println(run("aab"))
}

//分割回文串
/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。
s ="aab"  ["a","a","b"],["aa","b"]
*/
func run(s string) [][]string {
	res := [][]string{}
	dfs([]string{}, 0, &res, s)
	return res
}

func dfs(temp []string, start int, res *[][]string, s string) {
	if start == len(s) {
		t := append([]string(nil), temp...)
		*res = append(*res, t)
		return
	}

	for i := start; i < len(s); i++ {
		if valid(s, start, i) {
			//当前前段满足回文，切割判断下个序列是否连续包含回文子串
			temp = append(temp, s[start:i+1])
			dfs(temp, i+1, res, s)
			//剔除最后的单独字符串
			temp = temp[:len(temp)-1]
		}
	}
}

func valid(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

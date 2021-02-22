package main

import "fmt"

//给定一个 haystack 字符串和一个 needle 字符串，
//在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。
//如果不存在，则返回 -1
func main() {
	fmt.Println(strStr("asddasddd", "ddd"))
}

func strStr(haystack string, needle string) int {
	i, j := 0, 0
	for i = 0; i <= len(haystack)-len(needle); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

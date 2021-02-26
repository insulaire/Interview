package main

func main() {

}

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列
func checkInclusion(s1 string, s2 string) bool {
	need := map[byte]int{}
	win := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right, match := 0, 0, 0
	for right < len(s2) {
		//扩大窗口
		c := s2[right]
		right++
		//检测到目标元素 加入检测窗口
		if need[c] != 0 {
			win[c]++
			//检测值+1
			if need[c] == win[c] {
				match++
			}
		}

		//窗口大小超过或等于检测量 开始检测
		for right-left >= len(s1) {
			if match == len(need) {
				return true
			}
			//缩小窗口
			d := s2[left]
			left++
			//检测到目标元素
			if need[d] != 0 {
				//滑动后检测值-1
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}
	return false
}

package main

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

示例 4：
输入： dvdf
输出： 3
*/
//pwwkew
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := make(map[int32]int)
	lenMap := 0
	start := 0
	for index, val := range s {
		if v, ok := result[val]; ok {
			if v >= start {
				start = v + 1
			}
		}
		len := index - start + 1
		if lenMap < len {
			lenMap = len
		}
		result[val] = index
	}
	return lenMap
}

// func main() {
// 	s := "abcabcbb"
// 	fmt.Println("len s", len(s))
// 	res := lengthOfLongestSubstring(s)
// 	fmt.Println(res)
// }

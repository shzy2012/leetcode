package main

import "fmt"

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (31.90%)
 * Likes:    2677
 * Dislikes: 0
 * Total Accepted:    265.6K
 * Total Submissions: 831.1K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	if len(s) <= 0 {
		return 0
	}

	max := 0
	curlen := 0
	tmpStr := ""
	for i := 0; i < len(s); i++ {
		for index := 0; index < len(tmpStr); index++ {
			if s[i] == tmpStr[index] {
				tmpStr = tmpStr[index+1:]
				index = index - 1
				break
			}
		}

		tmpStr = tmpStr + string(s[i])
		curlen = len(tmpStr)
		if curlen > max {
			max = curlen
		}
	}

	return max
}

func lengthOfLongestSubstring2(s string) int {
	var i, j, k, max int
	for j = 0; j < len(s); j++ {
		for k = i; k < j; k++ {
			if s[k] == s[j] {
				i = k + 1
				break
			}
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}

	return max
}

func main() {

	max := lengthOfLongestSubstring("pwwkew")
	fmt.Println(max)
}

package leetcode

import "strings"

/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	rs := 0
	str := ""
	for _, c := range s {
		i := strings.IndexRune(str, c)
		if i >= 0 {
			str = str[i+1:] + string(c)
		} else {
			str += string(c)
		}
		if len(str) > rs {
			rs = len(str)
		}
	}
	return rs
}

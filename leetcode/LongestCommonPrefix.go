package leetcode

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) > 0 {
		minStr := strs[0]
		for _, str := range strs {
			if len(str) < len(minStr) {
				minStr = str
			}
		}
		rs := ""
		for i, c := range minStr {
			for _, str := range strs {
				if string(c) != string(str[i]) {
					return rs
				}
			}
			rs += string(c)
		}

		return rs
	} else {
		return ""
	}
}

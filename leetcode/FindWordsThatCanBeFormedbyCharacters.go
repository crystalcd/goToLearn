package leetcode

import (
	"fmt"
	"strings"
)

/**
You are given an array of strings words and a string chars.

A string is good if it can be formed by characters from chars (each character can only be used once).

Return the sum of lengths of all good strings in words.



Example 1:

Input: words = ["cat","bt","hat","tree"], chars = "atach"
Output: 6
Explanation:
The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
Example 2:

Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
Output: 10
Explanation:
The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.


Note:

1 <= words.length <= 1000
1 <= words[i].length, chars.length <= 100
All strings contain lowercase English letters only.

遍历每个字符串，每个字符串中的某个字符包含于chars里面就从chars中去掉这个字符再次校验知道结束判断是否是个good string
*/
func CountCharacters(words []string, chars string) int {
	rs := ""
	for _, str := range words {
		has := true
		stage := chars
		for _, s := range str {
			if strings.ContainsRune(stage, s) {
				i := strings.IndexRune(stage, s)
				stage = stage[:i] + stage[i+1:]
			} else {
				has = false
				break
			}
		}
		if has {
			rs += str
		}
	}
	fmt.Println(rs)
	return len(rs)
}

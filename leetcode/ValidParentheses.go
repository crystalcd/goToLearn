package leetcode

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

/**
采用一个数组模仿栈
一开始先入栈一个字符，后面字符与前面字符比较如果对应就弹出数组最后一个，不对应就插入数组尾部后面一直操作，最后数组长度为0则正确
*/
func IsValid(s string) bool {
	strMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	arr := make([]int32, 0)
	for _, c := range s {
		if len(arr) == 0 {
			arr = append(arr, c)
		} else {
			if string(c) != strMap[string(arr[len(arr)-1])] {
			} else {
				arr = arr[:len(arr)-1]
				continue
			}
			arr = append(arr, c)
		}
	}
	if len(arr) == 0 {
		return true
	}
	return false
}

package leetcode

/**
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

/**
1221 先比较首位和末位，末位直接x%10,首位需要先知道除以多少个10,比较完首位和末位获取中间的数字重复以上操作
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	for x/div >= 10 {
		div *= 10
	}
	for x != 0 {
		left := x / div
		right := x % 10
		if left != right {
			return false
		}
		x = (x % div) / 10
		div /= 100
	}
	return true
}

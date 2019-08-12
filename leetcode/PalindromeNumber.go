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
121 第一位和最后一位121/100和121%10 在比较2
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

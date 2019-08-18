package leetcode

/**
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: 3
Output: "III"
Example 2:

Input: 4
Output: "IV"
Example 3:

Input: 9
Output: "IX"
Example 4:

Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

1. 对num每一位去商 用罗马数字替换商
2. 根据每一位根据商的情况选择不同的罗马数字，比如小于4如300 CCC，等于4如400 CD，大于4小于9如800 DCCC，等于9如900 CM
3. 对num取余在用余数重复以上
*/
func intToRoman(num int) string {
	rs := ""
	strs := [...]string{
		"M", "D", "C", "L", "X", "V", "I",
	}
	vals := [...]int{
		1000, 500, 100, 50, 10, 5, 1,
	}
	for i := 0; i < len(vals); i += 2 {
		x := num / vals[i]
		if x < 4 {
			for j := 0; j < x; j++ {
				rs += strs[i]
			}
		} else if x == 4 {
			rs += strs[i] + strs[i-1]
		} else if x > 4 && x < 9 {
			rs += strs[i-1]
			for j := 5; j < x; j++ {
				rs += strs[i]
			}
		} else if x == 9 {
			rs += strs[i] + strs[i-2]
		}
		num %= vals[i]
	}
	return rs
}

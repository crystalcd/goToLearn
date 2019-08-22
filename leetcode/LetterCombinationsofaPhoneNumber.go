package leetcode

/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/
func LetterCombinations(digits string) []string {
	rs := make([]string, 0)
	if len(digits) == 0 {
		return rs
	}
	strMap := make(map[string]string)
	strMap["2"] = "abc"
	strMap["3"] = "def"
	strMap["4"] = "ghi"
	strMap["5"] = "jkl"
	strMap["6"] = "mno"
	strMap["7"] = "pqrs"
	strMap["8"] = "tuv"
	strMap["9"] = "wxyz"

	dfs(digits, strMap, 0, "", &rs)
	return rs
}

func dfs(digits string, dict map[string]string, level int, curStr string, rs *[]string) {
	if level == len(digits) {
		*rs = append(*rs, curStr)
		return
	}
	for _, c := range dict[string(digits[level])] {
		dfs(digits, dict, level+1, curStr+string(c), rs)
	}
}

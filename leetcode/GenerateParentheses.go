package leetcode

/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
/**
递归深度优先
*/
func GenerateParenthesis(n int) []string {
	rs := make([]string, 0)
	gendfs(n, n, "", &rs)
	return rs
}

func gendfs(l, r int, str string, rs *[]string) {
	if l > r {
		return
	}
	if l > 0 {
		gendfs(l-1, r, str+"(", rs)
	}
	if r > 0 {
		gendfs(l, r-1, str+")", rs)
	}
	if l == 0 && r == 0 {
		*rs = append(*rs, str)
	}
}

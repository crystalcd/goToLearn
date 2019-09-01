package leetcode

import "math"

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
func maxSubArray(nums []int) int {
	rs, curSum := math.MinInt32, 0
	for _, num := range nums {
		curSum = max(curSum+num, num)
		rs = max(rs, curSum)
	}
	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

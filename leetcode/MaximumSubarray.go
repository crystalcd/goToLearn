package leetcode

import "math"

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/
/**
O（n）
讲当前数和前面的数的和相比如果当前数比前数和大则重新一当前数为起点，最后找出最大值
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

/**
分治法
重要的是求解左中右以及包含中间的最大值
*/
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return divide(nums, 0, len(nums)-1)
}

func divide(nums []int, l, r int) int {
	if l >= r {
		return nums[l]
	}
	mid := (l + r) / 2
	mmax := nums[mid]
	lmax := divide(nums, l, mid-1)
	rmax := divide(nums, mid+1, r)
	tmp := mmax
	for i := mid - 1; i >= l; i-- {
		tmp += nums[i]
		mmax = max(mmax, tmp)
	}
	tmp = mmax
	for i := mid + 1; i <= r; i++ {
		tmp += nums[i]
		mmax = max(mmax, tmp)
	}
	return max(mmax, max(lmax, rmax))
}

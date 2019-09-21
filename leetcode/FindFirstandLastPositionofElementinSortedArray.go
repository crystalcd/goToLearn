package leetcode

/**
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	rsIndex := -1
	for l <= r {
		m := (r + l) / 2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			rsIndex = m
			break
		}
	}
	rs := make([]int, 0)
	if rsIndex == -1 {
		rs = append(rs, -1)
		rs = append(rs, -1)
		return rs
	} else {
		low := rsIndex
		for low > -1 && nums[low] == target {
			low--
		}
		rs = append(rs, low+1)
		high := rsIndex
		for high < len(nums) && nums[high] == target {
			high++
		}
		rs = append(rs, high-1)
		return rs
	}
}

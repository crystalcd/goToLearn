package leetcode

import (
	"fmt"
	"math"
)

/**
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func ThreeSumClosest(nums []int, target int) int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && key < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
	fmt.Println(nums)
	near := math.MaxUint32
	var x, y, z int
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			abs := nums[i] + nums[l] + nums[r] - target
			if abs > 0 {
				if abs < near {
					near = abs
					x, y, z = i, l, r
				}
				r--
			} else if abs < 0 {
				if -abs < near {
					near = -abs
					x, y, z = i, l, r
				}
				l++
			} else {
				l++
				r--
				return target
			}
		}
	}
	return nums[x] + nums[y] + nums[z]
}

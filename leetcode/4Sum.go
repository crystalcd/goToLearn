package leetcode

import (
	"fmt"
	"strconv"
)

/**
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func FourSum(nums []int, target int) [][]int {
	rs := make([][]int, 0)
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
	rsStr := make([]string, 0)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			l, r := j+1, len(nums)-1
			for l < r {
				curSum := nums[i] + nums[j] + nums[l] + nums[r]
				curArr := make([]int, 0)
				if curSum > target {
					r--
				} else if curSum < target {
					l++
				} else {
					curStr := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])
					if !isHasStr(&rsStr, curStr) {
						rsStr = append(rsStr, curStr)
						curArr = append(curArr, nums[i])
						curArr = append(curArr, nums[j])
						curArr = append(curArr, nums[l])
						curArr = append(curArr, nums[r])
						rs = append(rs, curArr)
					}
					rsStr = append(rsStr, curStr)
					l++
					r--
				}
			}
		}
	}
	return rs
}
func isHasStr(strs *[]string, target string) bool {
	for _, str := range *strs {
		if str == target {
			return true
		}
	}
	return false
}

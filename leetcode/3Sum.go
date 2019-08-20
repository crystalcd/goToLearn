package leetcode

import (
	"fmt"
	"strconv"
)

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
// 失败超时
func threeSum(nums []int) [][]int {
	rs := make([][]int, 0)
	strMap := make(map[string]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					arr := make([]int, 0)
					arr = append(arr, nums[i])
					arr = append(arr, nums[j])
					arr = append(arr, nums[k])
					for t := 0; t < len(arr); t++ {
						for z := t + 1; z < len(arr); z++ {
							if arr[z] < arr[t] {
								arr[z], arr[t] = arr[t], arr[z]
							}
						}
					}
					key := ""
					for _, s := range arr {
						key += strconv.Itoa(s)
					}
					fmt.Println(key)
					if strMap[key] == 0 {
						rs = append(rs, arr)
						strMap[key]++
					}
				}
			}
		}
	}
	return rs
}

/**
对原数组进行排序，然后开始遍历排序后的数组，这里注意不是遍历到最后一个停止，
而是到倒数第三个就可以了。这里可以先做个剪枝优化，就是当遍历到正数的时候就 break，
为啥呢，因为数组现在是有序的了，如果第一个要 fix 的数就是正数了，则后面的数字就都是正数，
就永远不会出现和为0的情况了。然后还要加上重复就跳过的处理，处理方法是从第二个数开始，
如果和前面的数字相等，就跳过，因为不想把相同的数字fix两次。对于遍历到的数，
用0减去这个 fix 的数得到一个 target，然后只需要再之后找到两个数之和等于 target 即可。
用两个指针分别指向 fix 数字之后开始的数组首尾两个数，如果两个数和正好为 target,左指针++右指针--，
则将这两个数和 fix 的数一起存入结果中。然后就是跳过重复数字的步骤了，
两个指针都需要检测重复数字。如果两数之和小于 target，则将左边那个指针i右移一位，
使得指向的数字增大一些。同理，如果两数之和大于 target，则将右边那个指针j左移一位，
使得指向的数字减小一些，
*/
func threeSum1(nums []int) [][]int {
	rs := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	fmt.Println(nums)
	outMap := make(map[string]int)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		target := 0 - nums[i]
		l, r := i+1, len(nums)-1

		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				arr := make([]int, 0)
				arr = append(arr, nums[i])
				arr = append(arr, nums[l])
				arr = append(arr, nums[r])
				key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])
				if outMap[key] == 0 {
					rs = append(rs, arr)
					outMap[key]++
				}
				l++
				r--
			}
		}
	}
	return rs
}

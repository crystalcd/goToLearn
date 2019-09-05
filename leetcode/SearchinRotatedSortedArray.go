package leetcode

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
0　　1　　2　　 4　　5　　6　　7

7　　0　　1　　 2　　4　　5　　6

6　　7　　0　　 1　　2　　4　　5

5　　6　　7　　 0　　1　　2　　4

4　　5　　6　　7　　0　　1　　2

2　　4　　5　　6　　7　　0　　1

1　　2　　4　　5　　6　　7　　0

二分搜索法的关键在于获得了中间数后，
判断下面要搜索左半段还是右半段，
我们观察上面红色的数字都是升序的，
由此我们可以观察出规律，如果中间的数小于最右边的数，
则右半段是有序的，若中间数大于最右边数，则
左半段是有序的，我们只要在有序的半段里用首尾
两个数组来判断目标值是否在这一区域内，这样就可以确定保留哪半边了，
代码如下：
*/
func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[r] {
			if target <= nums[r] && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

	}
	return -1
}

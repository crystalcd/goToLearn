package leetcode

import "fmt"

/**
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
/**
通过观察原数组可以发现，
如果从末尾往前看，数字逐渐变大，
到了2时才减小的，
然后我们再从后往前找第一个比2大的数字，是3，那么我们交换2和3，再把此时3后面的所有数字转置一下即可
1　　2　　7　　4　　3　　1

1　　2　　7　　4　　3　　1

1　　3　　7　　4　　2　　1

1　　3　　1　　2　　4　　7


*/
func NextPermutation(nums []int) {
	if len(nums) == 1 {
		fmt.Println(nums)
		return
	}
	index := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i - 1
			break
		}
	}
	next := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[index] {
			next = i
			break
		}
	}
	nums[index], nums[next] = nums[next], nums[index]
	newSlice := nums[index+1:]
	if index == next {
		newSlice = nums[index:]
	}

	for i := 0; i <= (len(newSlice)-1)/2; i++ {
		newSlice[i], newSlice[len(newSlice)-1-i] = newSlice[len(newSlice)-1-i], newSlice[i]
	}
	fmt.Println(nums)
}

func reverse(nums []int) {

}

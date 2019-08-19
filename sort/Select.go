package sort

/**
选择排序，每次fix一位，选出fix后最小的和fix替换
*/
func SelectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

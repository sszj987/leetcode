package num35

import (
	"fmt"
	"testing"
)

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func Test(t *testing.T) {
	if 1 == searchInsert([]int{1}, 2) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Error")
	}
}

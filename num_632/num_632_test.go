package num_632

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	result := smallestRange([][]int{[]int{4,10,15,24,26}, []int{0,9,12,20}, []int{5,18,22,30}})
	fmt.Println(result)
}


func smallestRange(nums [][]int) []int {
	sum := []int{}
	mmap := make(map[int]int)
	nextmap := make(map[int]int)
	leftnum, rightnum := 100000, 0
	for i:=0;i<len(nums);i++{
		nextmap[i] = 0
		for j :=0;j<len(nums[i]);j++{
			sum = append(sum, nums[i][j])
			mmap[nums[i][j]] = i
		}
		if len(nums[i]) != 0 {
			leftnum = min(nums[i][0], leftnum)
			rightnum = max(nums[i][0], rightnum)
		}
	}


	sort.Ints(sum)
	left := find(sum, leftnum)
	right := find(sum, rightnum)

	result := make([]int, 2)
	result[0] = left
	result[1] = right
	for right < len(sum) {
		i := mmap[sum[left]]
		nextmap[i] = nextmap[i]+1
		next := nextmap[i]
		if next >= len(nums[i]){
			break
		}

		if sum[next] > right {
			left++
			right = next
		} else if sum[next] <= right {
			left++
		}

		if sum[right] - sum[left] < result[1] - result[0] {
			result[0] = left
			result[1] = right
		}
	}

	return result
}

func find (slice []int, tar int) int {
	for i,v := range slice {
		if v == tar {
			return i
		}
	}
	return 0
}
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
package num_60

import (
	"fmt"
	"strconv"
	"testing"
)

func getPermutation(n int, k int) string {
	dp := make([]int, n+1)
	dp[1] = 1
	start := n-1
	for i:=2;i<=n;i++{
		dp[i] = i*dp[i-1]
	}

	isUsed := make([]int, n+1)
	var output string
	for k != 0 && start > 0 {
		var num int
		if k%dp[start] == 0 {
			num = k/dp[start]
		} else{
			num = k/dp[start] + 1
		}

		if k%dp[start] == 0{
			k = dp[start]
		} else {
			k = k%dp[start]
		}
		start--
		for i :=1;i<=n;i++{
			if isUsed[i] == 0 {
				num--
			}
			if num == 0 {
				isUsed[i] = 1
				output += strconv.Itoa(i)
				break
			}
		}
	}

	for len(output) < n {
		for i :=1;i<=n;i++{
			if isUsed[i] == 0 {
				isUsed[i] = 1
				output += strconv.Itoa(i)
			}
		}
	}
	return output
}

func TestNum60(t *testing.T) {
	output := getPermutation(3,1)
	fmt.Println(output == "123")
	output = getPermutation(3,2)
	fmt.Println(output == "132")
	output = getPermutation(3,3)
	fmt.Println(output == "213")
	output = getPermutation(3,4)
	fmt.Println(output == "231")
	output = getPermutation(3,5)
	fmt.Println(output == "312")
	output = getPermutation(3,6)
	fmt.Println(output == "321")

	output = getPermutation(1,1)
	fmt.Println(output == "1")
}
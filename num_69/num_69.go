package num69

func mySqrt(x int) int {
	start := 0
	end := x
	for start < end {
		mid := (start + end + 1) >> 1
		if mid*mid > x {
			end = mid - 1
		} else {
			start = mid
		}
	}
	return start
}

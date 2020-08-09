package num_93

import (
	"strconv"
	"testing"
)


func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) >  4 * 3 {
		return []string{}
	}
	result := []string{}
	result = append(result, "")
	result = restoreIpAddressesProc(s, result, 3)
	return result
}

func restoreIpAddressesProc(s string, input []string, step int) []string{
	lens := len(s)

	output := []string{}
	for i:=1;i<4;i++{
		if lens - i < step || lens - i >  step * 3 {
			continue
		}
		str := s[:i]
		if str[0] == '0' && len(str) != 1{
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		if i == 3 && num > 255 {
			continue
		}
		temp := s[:i]
		if step != 0 {
			temp += "."
		}

		tempSlice := []string{}
		for i, _ := range(input) {
			tempSlice = append(tempSlice, input[i] + temp)
		}
		if step != 0 {
			tempSlice = restoreIpAddressesProc(s[i:], tempSlice, step - 1)
		}
		output = append(output, tempSlice...)
	}
	return output
}

func Test(t *testing.T) {
	restoreIpAddresses("25525511135")
	restoreIpAddresses("010010")
}
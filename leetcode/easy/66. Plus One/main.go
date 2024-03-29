package main

import "fmt"

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1
	if digits[lastIndex] < 9 {
		digits[lastIndex]++
	} else {
		for i := lastIndex; i >= 0; i-- {
			digits[i] = 0

			if i > 0 && digits[i-1] < 9 {
				digits[i-1]++
				break
			}
		}

		if digits[0] == 0 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}

func main() {
	//digits := []int{9, 9, 9, 9}
	digits := []int{1, 8, 9}

	res := plusOne(digits)
	fmt.Println(res)
}

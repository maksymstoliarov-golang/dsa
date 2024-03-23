package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var last *int

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if last == nil {
			last = &val
		} else if *last == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else if *last != val {
			*last = val
		}
	}

	return len(nums)
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	//fmt.Println(nums)
	k := removeDuplicates(nums)

	fmt.Println(k)
	fmt.Println(nums)
}

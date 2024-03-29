package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start := 0
	lastIndex := len(nums) - 1
	end := lastIndex

	for start <= end {
		middle := (start + end) / 2

		if target < nums[middle] {
			end = middle - 1
		} else if target > nums[middle] {
			start = middle + 1
		} else {
			return middle
		}
	}

	if target < nums[0] {
		return 0
	}

	if target > nums[lastIndex] {
		return lastIndex + 1
	}

	return ((start + end) / 2) + 1
}

func main() {
	nums := []int{2, 6, 7, 9, 15}

	res := searchInsert(nums, 99)

	fmt.Println(res)
}

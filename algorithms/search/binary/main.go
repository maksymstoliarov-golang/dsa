package main

import "fmt"

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		middle := (start + end) / 2

		if nums[middle] > target {
			end = middle - 1
		} else if nums[middle] < target {
			start = middle + 1
		} else {
			return middle
		}
	}

	return -1
}

func main() {
	nums := []int{1, 3, 4, 7, 15, 22, 150}
	target := 150
	fmt.Println(search(nums, target))
}

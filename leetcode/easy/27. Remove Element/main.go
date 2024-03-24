package main

import "fmt"

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		v := nums[i]

		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

func main() {
	nums := []int{4, 5, 2, 1, 2, 3, 5, 6, 7, 9, 9, 76, 6, 4, 3, 3, 323, 2, 2, 2, 5, 7}
	fmt.Println(nums)
	fmt.Println(len(nums))
	k := removeElement(nums, 2)
	//fmt.Println(nums)
	fmt.Println(k)
}

package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, val := range nums {
		neededVal := target - val

		v, ok := m[neededVal]
		if ok {
			return []int{v, i}
		}
		m[val] = i
	}

	return nil
}

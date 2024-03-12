package main

import "fmt"

func Sort(s []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				sorted = false
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}

func main() {
	fmt.Println(Sort([]int{5, 4, 23, 3, 2, 7, 8, 5, 4, 3, 0. - 22}))
}

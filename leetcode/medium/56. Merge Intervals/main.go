package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	l := 0
	m := false

	for r := 1; r < len(intervals); r++ {
		if intervals[l][1] >= intervals[r][0] {
			intervals[l][0] = min(intervals[l][0], intervals[r][0])
			intervals[l][1] = max(intervals[l][1], intervals[r][1])

			intervals = append(intervals[:r], intervals[r+1:]...)
			m = true
		}

		if m == true {
			r--
			m = false
		} else {
			l++
		}
	}

	return intervals
}

func main() {
	var intervals [][]int
	intervals = [][]int{{1, 4}, {3, 5}, {1, 9}, {9, 10}, {11, 12}, {11, 15}}

	fmt.Println(intervals)

	//intervals = [][]int{{1, 4}, {2, 3}}

	//intervals = [][]int{{1, 4}, {0, 1}}

	//intervals = [][]int{{0, 0}, {1, 4}}

	res := merge(intervals)

	fmt.Println(res)
}

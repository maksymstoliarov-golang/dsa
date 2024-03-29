package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.Join(strings.Fields(strings.TrimSpace(s)), " "), " ")

	return len(split[len(split)-1])
}

func lengthOfLastWord2(s string) int {
	i, l := len(s)-1, 0

	for s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		i--
		l++
	}

	return l
}

func main() {
	s := "  Hello    World      "

	res := lengthOfLastWord2(s)

	fmt.Println(res)
}

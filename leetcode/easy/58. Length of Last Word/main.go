package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.Join(strings.Fields(strings.TrimSpace(s)), " "), " ")

	return len(split[len(split)-1])
}

func main() {
	s := "  Hello    World      "

	res := lengthOfLastWord(s)

	fmt.Println(res)
}

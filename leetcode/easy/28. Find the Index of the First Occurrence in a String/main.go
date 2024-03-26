package main

import (
	"fmt"
	"reflect"
)

func strStr(haystack string, needle string) int {
	if needle == "" || haystack == "" {
		return -1
	}

	hs := []byte(haystack)
	nee := []byte(needle)
	neeLen := len(needle)

	for i, ch := range hs {
		if ch == nee[0] {
			if len(hs[i:]) < neeLen {
				return -1
			}

			possible := hs[i : i+neeLen]

			if reflect.DeepEqual(possible, nee) {
				return i
			}
		} else if len(hs[i:]) < neeLen {
			return -1
		}
	}

	return -1
}

func main() {
	// mississippi
	haystack := "mississippi"

	i := strStr(haystack, "issin")
	fmt.Println(i)
}

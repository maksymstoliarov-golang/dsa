package main

import "fmt"

// (){}[] valid
// (] invalid

type Node struct {
	val  string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) pop() string {
	val := l.head.val
	l.head = l.head.next
	return val
}

func (l *LinkedList) push(val string) {
	l.head = &Node{val: val, next: l.head}
}

func isValid(s string) bool {
	//if string contains 1,3,5 characters, then obviously it is not valid
	if len(s)%2 == 1 {
		return false
	}

	m := map[string]string{")": "(", "]": "[", "}": "{"}

	//create empty linked list
	l := new(LinkedList)

	//check each character of string | if it is opening, then add to linked list, then add | if closing, compare with head and pop
	for i := 0; i < len(s); i++ {
		ch := fmt.Sprintf("%c", s[i])
		if isOpening(ch) {
			l.push(ch)
		} else {
			if l.head == nil {
				return false
			}

			for l.head != nil && m[ch] != l.pop() {
				return false
			}
		}
	}

	if l.head != nil {
		return false
	}

	return true
}

func isOpening(s string) bool {
	return s == "{" || s == "[" || s == "("
}

func main() {
	s := "(({}))[]"

	r := isValid(s)

	fmt.Println(s)
	fmt.Println(r)
}

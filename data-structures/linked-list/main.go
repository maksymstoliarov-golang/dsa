package main

import "fmt"

type Item struct {
	val  int
	next *Item
}

type List struct {
	head *Item
}

func (l *List) push(val int) {
	l.head = &Item{val: val, next: l.head}
}

func (l *List) remove(val int) {
	if l.head == nil {
		return
	}

	if l.head.val == val {
		l.head = l.head.next
		return
	}

	head := l.head
	for head.next != nil {
		if head.next.val == val {
			head.next = head.next.next
			return
		}

		head = head.next
	}
}

func (l *List) removeAll(val int) {
	if l.head == nil {
		return
	}

	for l.head != nil && l.head.val == val {
		l.head = l.head.next
	}

	head := l.head
	for head.next != nil {
		if head.next.val == val {
			head.next = head.next.next
		} else {
			head = head.next
		}

	}
}

func (l *List) find(val int) *Item {
	head := l.head
	for head != nil {
		if head.val == val {
			return head
		}
		head = head.next
	}
	return nil
}

func (l *List) print() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	head := l.head
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}

func main() {
	l := &List{head: &Item{val: 2, next: &Item{val: 3, next: &Item{val: 2, next: &Item{val: 4}}}}}

	//l.print()
	//
	//l.push(4)
	//l.push(5)
	//
	//l.print()
	//
	//f := l.find(5)
	//fmt.Println(f)

	//l.remove(4)
	//l.remove(5)
	//l.remove(2)
	//l.remove(2)
	//l.remove(2)
	l.removeAll(2)
	l.print()
}

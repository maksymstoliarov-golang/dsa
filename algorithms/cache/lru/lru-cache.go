package main

import (
	"fmt"
)

type Item struct {
	key  int
	val  int
	prev *Item
	next *Item
}

type List struct {
	head *Item
	tail *Item
}

func (l *List) put(key, val int) *Item {
	oldHead := l.head
	item := &Item{key: key, val: val, next: oldHead}
	if oldHead != nil {
		oldHead.prev = item
	}

	l.head = item

	if l.tail == nil {
		l.tail = item
	}

	return item
}

func (l *List) remove(item *Item) {
	//update prev
	if item.prev != nil {
		item.prev.next = item.next
	}

	//update next
	if item.next != nil {
		item.next.prev = item.prev
	}

	//if tail
	if l.tail == item {
		l.tail = item.prev
	}

	//if head
	if l.head == item {
		l.head = item.next
	}
}

func (l *List) print() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	head := l.head
	for head != nil {
		//fmt.Println(head.key, "->", head.val)
		fmt.Println(head)
		head = head.next
	}
}

type LRUCache struct {
	list     *List
	data     map[int]*Item
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, data: make(map[int]*Item), list: new(List)}
}

func (c *LRUCache) Get(key int) int {
	item, ok := c.data[key]
	if !ok {
		return -1
	}

	//if v.prev != nil {
	//	v.prev.next = v.next
	//}
	//v.next = c.list.head
	//c.list.head = v
	if c.list.head.key == key {
		return item.val
	}

	if c.list.head == c.list.tail {
		return item.val
	}

	//if c.list.tail != nil && c.list.tail.key == key {
	//	c.list.tail = v.prev
	//}

	//update prev
	if item.prev != nil {
		item.prev.next = item.next
	}

	//update next
	if item.next != nil {
		item.next.prev = item.prev
	}

	oldHead := c.list.head

	item.next = oldHead
	oldHead.prev = item

	c.list.head = item
	//c.list.tail = oldHead
	//if oldHead.next == c.list.tail {
	//	c.list.tail = oldHead
	//}

	if c.list.tail == item {
		c.list.tail = item.prev
	}
	item.prev = nil

	return item.val
}

func (c *LRUCache) Put(key int, value int) {
	//if already exists
	if v, ok := c.data[key]; ok {
		//check if this is tail and set tail to prev
		//if c.list.tail == v && v.prev != nil {
		//	c.list.tail = v.prev
		//}

		//update value of existing node
		v.val = value

		c.list.remove(v)
		newItem := c.list.put(key, value)
		c.data[key] = newItem
		return
	}

	//check for capacity
	if len(c.data) >= c.capacity {
		//remove lru from list and map
		lruKey := c.list.tail.key
		//fmt.Println(lruKey)
		c.list.remove(c.list.tail)
		delete(c.data, lruKey)
	}

	//put new value to list and map
	newItem := c.list.put(key, value)
	c.data[key] = newItem

}

//"LRUCache 2",

//"put 1:1" null
//"put 2"2" null
//"get 1" 1
//"put 3:3" null
//"get 2" -1
//"put 4:4" null
//"get 1" -1
//"get 3" 3
//"get 4" 4

func main() {
	c := Constructor(2)
	c.Put(1, 1) // cache is {1=1}
	c.Put(2, 2) // cache is {1=1, 2=2}
	c.Get(1)    // return 1
	c.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	c.Get(2)    // returns -1 (not found)
	c.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	c.Get(1)    // return -1 (not found)
	c.Get(3)    // return 3
	c.Get(4)    // return 4
}

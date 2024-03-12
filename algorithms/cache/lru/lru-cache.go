package main

import "fmt"

type Item struct {
	key  int
	val  int
	prev *Item
	next *Item
}

type List struct {
	head *Item
	tail *Item //TODO: implement tail
}

func (l *List) put(key, val int) *Item {
	item := &Item{key: key, val: val, next: l.head}
	l.head = item
	return item
}

func (l *List) remove(item *Item) {
	item.prev.next = item.next
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

type LRUCache struct {
	list     *List
	data     map[int]*Item
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, data: make(map[int]*Item), list: new(List)}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.data[key]
	if !ok {
		return -1
	}

	if v.prev != nil {
		v.prev.next = v.next
	}
	v.next = c.list.head
	c.list.head = v

	return v.val
}

func (c *LRUCache) Put(key int, value int) {
	if v, ok := c.data[key]; ok {
		//update map value and move to head
		v.val = value
		c.list.remove(v)
		newItem := c.list.put(key, value)
		c.data[key] = newItem
	} else {
		//check for capacity
		if len(c.data) >= c.capacity {
			//remove lru from list and map
			lruKey := c.list.tail.key
			c.list.remove(c.list.tail)
			delete(c.data, lruKey)
		}

		//put new value to list and map
		newItem := c.list.put(key, value)
		c.data[key] = newItem
	}
}

func main() {
	//m := make(map[int]*Item)
	//list := new(List)
	//
	//m[3] = list.put(3)
	//
	//fmt.Println(m[3])

	cache := Constructor(3)
	fmt.Println(cache)

	cache.Put(2, 3)
	cache.Put(3, 4)

	fmt.Println(cache)
	cache.list.print()
}

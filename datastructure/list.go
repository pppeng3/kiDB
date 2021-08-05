package datastructure

import (
	"bytes"
)

type List struct {
	first *node
	last *node
	size int
}

type node struct {
	val []byte
	prev *node
	next *node
}

func NewList(vals ...[]byte) *List {
	list := List{}
	for _, v := range vals {
		list.Add(v)
	}
	return &list
}

func (list *List) Add(val []byte) {
	if list == nil {
		panic("list is nil")
	}
	n := &node {
		val: val,
	}
	if list.last == nil {
		list.first = n
		list.last = n
	} else {
		n.prev = list.last
		list.last.next = n
		list.last = n
	}
	list.size++
}

func (list *List) find(index int) (n *node) {
	if index < list.size / 2 {
		n = list.first
		for i := 0; i < index ; i++ {
			n = n.next
		}
	} else {
		n = list.last
		for i := list.size - 1; i > index; i-- {
			n = n.prev
		}
	}
	return n
}

func (list *List) Get(index int) (val []byte){
	if list == nil {
		panic("list is nil")
	}
	if index < 0 || index > list.size {
		panic("index out of bound")
	}
	return list.find(index).val
}

func (list *List) Set(index int, val []byte) bool {
	if list == nil {
		panic("list is nil")
	}
	if index < 0 || index >list.size {
		panic("index out of range")
	}
	n := list.find(index)
	n.val = val
	return true
}

func (list *List) Insert(index int, val []byte) {
	if list == nil {
		panic("list is nil")
	}
	if index < 0 || index > list.size {
		panic("index out of range")
	}

	if index == list.size {
		list.Add(val)
		return
	}

	pivot := list.find(index)
	n := &node{
		val: val,
		prev: pivot.prev,
		next: pivot,
	}
	if pivot.prev == nil {
		list.first = n
	} else {
		pivot.prev.next = n
	}
	pivot.prev = n
	list.size++
}

func (list *List) removeNode(n *node) {
	if n.prev == nil {
		list.first = n.next
	} else {
		n.prev.next = n.next
	}
	n.prev = nil
	n.next = nil
}

func (list *List) Remove(index int) bool {
	if list == nil {
		panic("list is nil")
	}
	if index < 0 || index >= list.size {
		panic("index out of bound")
	}

	n := list.find(index)
	list.removeNode(n)
	list.size--
	return true
}

func (list *List) RemoveLast() bool {
	if list == nil {
		panic("list is nil")
	}
	n := list.last
	list.removeNode(n)
	list.size--
	return true
}

func (list *List) RemoveAllByVal(val []byte) int {
	if list == nil {
		panic("list is nil")
	}
	n := list.first
	removed := 0
	var nextNode *node
	for n != nil {
		nextNode = n.next
		if bytes.Equal(n.val, val) {
			list.removeNode(n)
			removed++
		}
		n = nextNode
	}
	return removed
}

func (list *List) RemoveByVal(val []byte, count int) int {
	if list == nil {
		panic("list is nil")
	}
	n := list.first
	removed := 0
	var nextNode *node
	for n != nil {
		nextNode = n.next
		if bytes.Equal(n.val, val) {
			list.removeNode(n)
			removed++
		}
		if removed == count {
			break
		}
		n = nextNode
	}
	return removed
}

func (list *List) ReverseRemoveByVal(val []byte, count int) int {
	if list == nil {
		panic("list is nil")
	}
	n := list.last
	removed := 0
	var prevNode *node
	for n != nil {
		prevNode = n.prev
		if bytes.Equal(n.val, val) {
			list.removeNode(n)
			removed++
		}
		if removed == count {
			break
		}
		n = prevNode
	}
	return removed
}

func (list *List) Len() int {
	if list == nil {
		panic("list is nil")
	}
	return list.size
}

func (list *List) Contains(val []byte) bool {
	n := list.first
	for n != nil {
		if bytes.Equal(n.val, val) {
			return true
		}
		n = n.next
	}
	return false
}

func (list *List) Range(start, end int) []interface{} {
	if list == nil {
		panic("list is nil")
	}
	if start < 0 || start >= list.size || end < start || end > list.size {
		panic("index out of range")
	}

	sliceSize := end - start
	slice := make([]interface{}, sliceSize)
	n := list.first
	sliceIndex := 0
	for n != nil {
		if sliceIndex >= start && sliceIndex < end {
			slice[sliceIndex] = n.val
			sliceIndex++
		} else if sliceIndex > end {
			break
		}
		n = n.next
	}
	return slice
}


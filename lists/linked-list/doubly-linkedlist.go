package linked_list

import "github.com/Haevnen/data-structure/lists"

var _ lists.List[any] = (*DoublyLinkedList[any])(nil)

type DoublyLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

type node[T comparable] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func New[T comparable](values ...T) *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *DoublyLinkedList[T]) Add(values ...T) {
	for _, v := range values {
		newNode := &node[T]{value: v, prev: list.tail}
		if list.head == nil {
			list.head = newNode
			list.tail = newNode
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}
		list.size++
	}
}

func (list *DoublyLinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= list.size {
		var t T
		return t, false
	}

	var cur *node[T]
	if index < list.size/2 {
		// count from head
		cur = list.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
	} else {
		// count from tail
		cur = list.tail
		for i := list.size - 1; i > index; i-- {
			cur = cur.prev
		}
	}
	return cur.value, true
}

func (list *DoublyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= list.size {
		return
	}

	removedNode := list.head
	for i := 0; i < index; i++ {
		removedNode = removedNode.next
	}

	if list.size == 1 {
		list.head = nil
		list.tail = nil
		return
	}

	if removedNode == list.head {
		list.head = removedNode.next
	} else if removedNode == list.tail {
		list.tail = removedNode.prev
	} else {
		removedNode.prev.next = removedNode.next
		removedNode.next.prev = removedNode.prev
	}

	removedNode = nil
	list.size--
}

func (list *DoublyLinkedList[T]) Size() int {
	return list.size
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *DoublyLinkedList[T]) Set(index int, value T) bool {

}

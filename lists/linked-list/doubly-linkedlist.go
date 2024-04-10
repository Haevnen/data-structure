package linked_list

import "github.com/Haevnen/data-structure/lists"

var _ lists.List[any] = (*DoublyLinkedList[any])(nil)

type DoublyLinkedList[T comparable] struct {
}

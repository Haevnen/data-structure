package main

import (
	"fmt"

	linked_list "github.com/Haevnen/data-structure/lists/linked-list"
)

func main() {
	doublyLL := linked_list.New[int](45, 98, 2, 4)
	fmt.Println(doublyLL)

	fmt.Println(doublyLL.Values())
	fmt.Println(doublyLL.Get(0))
	fmt.Println(doublyLL.Get(2))

	doublyLL.Set(2, 99)
	fmt.Println(doublyLL)

	doublyLL.Sort(func(x, y int) int {
		if x > y {
			return -1
		}

		if x < y {
			return 1
		}

		return 0
	})
	fmt.Println(doublyLL)

	doublyLL.Remove(1)
	fmt.Println(doublyLL)

	doublyLL.Remove(2)
	fmt.Println(doublyLL)
	doublyLL.Remove(0)
	fmt.Println(doublyLL)
	doublyLL.Remove(0)
	fmt.Println(doublyLL)

	fmt.Println(doublyLL.IsEmpty())
	doublyLL.Remove(0)
	fmt.Println(doublyLL)
	doublyLL.Add(23)
	fmt.Println(doublyLL)
}

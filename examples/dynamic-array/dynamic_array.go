package main

import (
	"fmt"

	dynamic_array "github.com/Haevnen/data-structure/dynamic-array"
)

func main() {
	da := dynamic_array.New[int](1, 2, 3, 4)
	fmt.Printf("%v\n", da)
	da.Add(5)
	fmt.Printf("%v\n", da) // 1 2 3 4 5
	da.Sort(func(x, y int) int {
		if x > y {
			return -1
		}

		if x < y {
			return 1
		}

		return 0
	})
	fmt.Printf("%v\n", da)   // 5 4 3 2 1
	fmt.Println(da.Get(1))   // 4 true
	fmt.Println(da.Get(100)) // 0 false

	da.Remove(1) // 5 3 2 1
	fmt.Printf("%v\n", da)
	da.Remove(0)                         // 3 2 1
	fmt.Println("IsEmpty", da.IsEmpty()) // false

	fmt.Println(da.Values()) // 3 2 1
	da.Clear()               // empty
	fmt.Println(da)
}

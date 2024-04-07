package dynamic_array

import (
	"fmt"
	"slices"
	"strings"
)

// Comparable type support == operator (equality and inequality)
// Ordered type support both equality checks and comparison. E.g., <, <=, >, >=
type DynamicArray[T comparable] struct {
	elements []T
	size     int
	capacity int
}

// Public functions
func New[T comparable](values ...T) *DynamicArray[T] {
	da := &DynamicArray[T]{}

	if len(values) > 0 {
		da.Add(values...)
	}
	return da
}

func (da *DynamicArray[T]) Len() int {
	return da.size
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.size == 0
}

func (da *DynamicArray[T]) Get(index int) (T, bool) {
	if index < 0 || index >= da.size {
		var t T
		return t, false
	}
	return da.elements[index], true
}

func (da *DynamicArray[T]) Set(index int, value T) bool {
	if index < 0 || index >= da.size {
		return false
	}

	da.elements[index] = value
	return true
}

func (da *DynamicArray[T]) Add(values ...T) {
	da.expand(len(values))
	for _, v := range values {
		da.elements[da.size] = v
		da.size++
	}
}

func (da *DynamicArray[T]) Remove(index int) {
	if index < 0 || index >= da.size {
		return
	}

	// Shift all elements to the left
	for i := index + 1; i < da.size; i++ {
		da.elements[i-1] = da.elements[i]
	}
	da.size--
	da.shrink()
}

func (da *DynamicArray[T]) IndexOf(value T) int {
	if da.size > 0 {
		for idx, val := range da.elements {
			if val == value {
				return idx
			}
		}
	}
	return -1
}

func (da *DynamicArray[T]) Clear() {
	da.size = 0
	da.elements = make([]T, 0)
}

func (da *DynamicArray[T]) String() string {
	values := make([]string, da.size)
	for _, val := range da.elements {
		values = append(values, fmt.Sprintf("%v", val))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (da *DynamicArray[T]) Values() []T {
	values := make([]T, da.size)
	copy(values, da.elements)
	return values
}

func (da *DynamicArray[T]) Sort(comp func(x, y T) int) {
	if da.size < 2 {
		return
	}
	slices.SortFunc(da.elements, comp)
}

// Private functions
func (da *DynamicArray[T]) resize(cap int) {
	da.capacity = cap
	newElems := make([]T, cap)
	copy(newElems, da.elements)
	da.elements = newElems
}

func (da *DynamicArray[T]) expand(n int) {
	if da.size+n >= da.capacity {
		da.resize(da.capacity * 2)
	}
}

func (da *DynamicArray[T]) shrink() {
	if da.size <= int(da.capacity/4) {
		da.resize(da.size)
	}
}

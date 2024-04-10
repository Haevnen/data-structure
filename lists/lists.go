package lists

import "github.com/Haevnen/data-structure/containers"

type List[T comparable] interface {
	containers.Container[T]

	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Sort(comp func(x, y T) int)
	Set(index int, value T) bool
}

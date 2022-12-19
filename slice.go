package slices

import (
	"fmt"
)

type Slice[T any] struct {
	ts []T
}

func New[T any](ts []T) *Slice[T] {

	myTs := make([]T, len(ts))
	copy(myTs, ts)

	return &Slice[T]{
		ts: myTs,
	}
}

func (s Slice[T]) Elements() []T {
	return s.ts
}

func (s Slice[T]) String() string {
	return fmt.Sprint(s.ts)
}

func (s Slice[T]) Map(conv func(t T) interface{}) *Slice[interface{}] {
	return &Slice[interface{}]{
		ts: Map(conv, s.ts),
	}
}

func (s Slice[T]) Filter(filter func(t T) bool) *Slice[T] {
	return &Slice[T]{
		ts: inPlaceFilter(filter, s.ts),
	}
}

func (s Slice[T]) Reverse() *Slice[T] {
	return &Slice[T]{
		ts: Reverse(s.ts),
	}
}

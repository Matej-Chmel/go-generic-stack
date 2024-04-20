// Package for simple LIFO stack data structure.
package gostack

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// LIFO stack, Last In, First Out.
type Stack[T any] struct {
	data []T
}

// Returns the capacity of the underlying slice.
func (s *Stack[T]) Cap() int {
	return cap(s.data)
}

// Removes all items. Capacity remains unchanged.
func (s *Stack[T]) Clear() {
	s.ClearWithCapacity(s.Cap())
}

// Removes all items and sets capacity to c.
func (s *Stack[T]) ClearWithCapacity(c int) {
	s.data = make([]T, 0, c)
}

// Returns bool indicating if no items are in the stack.
func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

// Returns bool indicating if at least one item is in the stack.
func (s *Stack[T]) HasItems() bool {
	return len(s.data) > 0
}

func (s *Stack[T]) guard() error {
	if s.Empty() {
		return errors.New("cannot pop empty stack")
	}

	return nil
}

// Returns numnber of items in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// Removes the top item and returns it. If stack is empty, error is set.
func (s *Stack[T]) Pop() (T, error) {
	if err := s.guard(); err != nil {
		var none T
		return none, err
	}

	last := len(s.data) - 1
	top := s.data[last]
	s.data = s.data[:last]
	return top, nil
}

// Returns n top items or less.
func (s *Stack[T]) PopAvailable(n int) []T {
	sliceLen := min(s.Len(), n)
	return s.popSlice(sliceLen)
}

// Returns n top items or an error.
func (s *Stack[T]) PopExact(n int) ([]T, error) {
	if s.Len() < n {
		return make([]T, 0), errors.New("not enough elements")
	}

	return s.popSlice(n), nil
}

func (s *Stack[T]) popSlice(n int) []T {
	slice := make([]T, 0, n)

	for i := 0; i < n; i++ {
		top, _ := s.Pop()
		slice = append(slice, top)
	}

	return slice
}

// Adds an item at the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Adds multiple items at the top of the stack.
// The last item in values will be the new top item in the stack.
func (s *Stack[T]) PushMore(values ...T) {
	s.data = append(s.data, values...)
}

// Adds items from a slice at the top of the stack.
// The last item in the slice will be the new top item in the stack.
func (s *Stack[T]) PushSlice(values []T) {
	s.data = append(s.data, values...)
}

// Returns a string representation of the stack.
// Uses reflection to get the representation of every item.
func (s Stack[T]) String() string {
	if s.Empty() {
		return "[]"
	}

	builder := &strings.Builder{}
	builder.WriteRune('[')
	last := s.Len() - 1

	for i := 0; i < last; i++ {
		s.writeItem(builder, i, ", ")
	}

	s.writeItem(builder, last, "]")
	return builder.String()
}

// Attempts to update the capacity to n or returns an error.
func (s *Stack[T]) UpdateCapacity(n int) error {
	c := s.Cap()

	if n < c {
		return errors.New("cannot decrease capacity")
	}

	if n != c {
		data := make([]T, s.Len(), n)
		copy(data, s.data)
		s.data = data
	}

	return nil
}

// Returns a copy of the top item without removing it from the stack.
// If the stack is empty, an error is returned.
func (s *Stack[T]) Top() (T, error) {
	if err := s.guard(); err != nil {
		var none T
		return none, err
	}

	return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) writeItem(builder *strings.Builder, i int, sep string) {
	val := reflect.ValueOf(s.data[i])
	content := fmt.Sprintf("%v", val)
	builder.WriteString(content)
	builder.WriteString(sep)
}

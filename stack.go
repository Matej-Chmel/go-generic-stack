// Package for generic LIFO stack data structure
package genericstack

import "strings"

// Generic LIFO stack, Last In, First Out
type Stack[T any] struct {
	data []T
}

// Constructs a new stack
func New[T any]() Stack[T] {
	return Stack[T]{
		data: make([]T, 0, 0),
	}
}

// Returns the capacity of the the underlying slic
func (s *Stack[T]) Cap() int {
	return cap(s.data)
}

// Removes all items, capacity remains unchanged
func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}

// Removes all items and sets capacity to c
func (s *Stack[T]) ClearWithCap(c int) {
	if c == s.Cap() {
		s.Clear()
	} else {
		s.data = make([]T, 0, c)
	}
}

// Returns a flag indicating whether there are no items in the stack
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Format stack into a string according to Options
func (s *Stack[T]) Format(f *Format[T]) string {
	if f.Conversion == nil {
		f.Conversion = DefaultConversion[T]
	}

	if s.Empty() {
		return f.Start + f.End
	}

	var builder strings.Builder
	builder.WriteString(f.Start)

	if f.TopFirst {
		s.formatTopFirst(&builder, f)
	} else {
		s.formatBottomFirst(&builder, f)
	}

	builder.WriteString(f.End)
	return builder.String()
}

// Internal implementation of formatting items from bottom to top
func (s *Stack[T]) formatBottomFirst(builder *strings.Builder, f *Format[T]) {
	builder.WriteString(f.Conversion(&s.data[0]))
	length := s.Len()

	for i := 1; i < length; i++ {
		builder.WriteString(f.Sep)
		builder.WriteString(f.Conversion(&s.data[i]))
	}
}

// Internal implementation of formatting items from top to bottom
func (s *Stack[T]) formatTopFirst(builder *strings.Builder, f *Format[T]) {
	builder.WriteString(f.Conversion(s.TopPointer()))

	for i := s.lastIndex() - 1; i >= 0; i-- {
		builder.WriteString(f.Sep)
		builder.WriteString(f.Conversion(&s.data[i]))
	}
}

// Returns a flag indicating whether there is at least one item in the stack
func (s *Stack[T]) HasItems() bool {
	return s.Len() > 0
}

// Returns the index of the top item in the underlying slice
func (s *Stack[T]) lastIndex() int {
	return s.Len() - 1
}

// Returns numnber of items in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// Removes the top item.
// Panics if the stack is empty.
func (s *Stack[T]) Pop() {
	s.data = s.data[:s.lastIndex()]
}

// Removes the top item and returns a shallow copy of it.
// Panics if the stack is empty.
func (s *Stack[T]) PopAndReturn() T {
	res := s.data[s.lastIndex()]
	s.Pop()
	return res
}

// Adds an item as the new top of the stack.
func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Adds items on the stack in order.
// Last item becomes the new top.
func (s *Stack[T]) PushItems(items ...T) {
	s.PushSlice(items)
}

// Adds a slice of items on the stack in order.
// Last item becomes the new top.
func (s *Stack[T]) PushSlice(items []T) {
	for _, v := range items {
		s.Push(v)
	}
}

// Returns default string representation in form [top ... ...]
func (s Stack[T]) String() string {
	return s.Format(NewFormat[T]())
}

// Returns a shallow copy of the top item without removing it from the stack.
// Panics if the stack is empty.
func (s *Stack[T]) Top() T {
	return *s.TopPointer()
}

// Returns a pointer to the top item without removing it from the stack.
// Panics if the stack is empty.
func (s *Stack[T]) TopPointer() *T {
	return &s.data[s.lastIndex()]
}

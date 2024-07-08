package genericstack

import "fmt"

// Formatting options for a Stack[T]
type Format[T any] struct {
	// Function that converts each item of the stack to a string
	Conversion func(*T) string
	// Symbol to write after the last item
	End string
	// Symbol to write between two items
	Sep string
	// Symbol to write before the first item
	Start string
	// If true, items are written from top to bottom,
	// if false, the direction is reversed
	TopFirst bool
}

// Constructs new formatting options
func NewFormat[T any]() *Format[T] {
	return &Format[T]{
		Conversion: nil,
		End:        DefaultEnd,
		Sep:        DefaultSep,
		Start:      DefaultStart,
		TopFirst:   DefaultTopFirst,
	}
}

// Default conversion function converts each item to its default Go format
func DefaultConversion[T any](val *T) string {
	return fmt.Sprintf("%v", *val)
}

const (
	// Default symbol to write after the last item
	DefaultEnd string = "]"
	// Default symbol to write between two items
	DefaultSep string = " "
	// Default symbol to write before the first item
	DefaultStart string = "["
	// Default direction of formatting
	DefaultTopFirst bool = true
)

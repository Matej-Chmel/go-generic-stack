package gostack_test

import (
	"runtime"
	"testing"

	gostack "github.com/Matej-Chmel/go-stack"
)

func check[T comparable](a, b T, t *testing.T) {
	if a != b {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("Error at line %d, \"%v\" != \"%v\"", line, a, b)
	}
}

func checkPop[T comparable](stack *gostack.Stack[T], expected T, t *testing.T) {
	top, err := stack.Pop()
	check(top, expected, t)
	check(err, nil, t)
}

func checkTop[T comparable](stack *gostack.Stack[T], expected T, t *testing.T) {
	top, err := stack.Top()
	check(top, expected, t)
	check(err, nil, t)
}

func runTest[T comparable](data []T, t *testing.T) {
	stack := &gostack.Stack[T]{}
	stack.PushSlice(data)

	check(stack.Empty(), false, t)
	check(stack.HasItems(), true, t)
	check(stack.Len(), len(data), t)

	stack.UpdateCapacity(stack.Len() * 2)
	check(stack.Cap(), len(data)*2, t)
	check(stack.Len(), len(data), t)

	for i := len(data) - 1; i >= 0; i-- {
		checkTop(stack, data[i], t)
		checkPop(stack, data[i], t)
	}

	top, err := stack.Pop()
	var none T
	check(top, none, t)
	check(err.Error(), "cannot pop empty stack", t)

	check(stack.Empty(), true, t)
	check(stack.Len(), 0, t)
}

func TestComplex64(t *testing.T) {
	runTest([]complex64{1 + 1i, 2 + 1i, -3 + 2.4i, -5.2 + -1.01i}, t)
}

func TestInt32(t *testing.T) {
	runTest([]int32{23, 37, 420, -1190}, t)
}

func TestInt32Simple(t *testing.T) {
	stack := &gostack.Stack[int32]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	top, err := stack.Top()
	check(3, top, t)
	check(nil, err, t)

	top, err = stack.Pop()
	check(3, top, t)
	check(nil, err, t)

	check("[1, 2]", stack.String(), t)
	check(stack.Len(), 2, t)

	stack.ClearWithCapacity(8)
	check(stack.Empty(), true, t)
	check(stack.Cap(), 8, t)
}

func TestString(t *testing.T) {
	runTest([]string{"Hello, World!", "func main() {}", "..."}, t)
}

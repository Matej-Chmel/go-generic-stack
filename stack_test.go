package genericstack_test

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"testing"

	gs "github.com/Matej-Chmel/go-generic-stack"
)

// Type for testing custom struct types
type example struct {
	flag bool
	val  int
}

// Wrapper around original test type
type tester struct {
	failed bool
	*testing.T
}

// Constructs new Tester and a stack of type T
func newTester[T any](t *testing.T) (tester, gs.Stack[T]) {
	return tester{failed: false, T: t}, gs.New[T]()
}

// Check whether all pairs from values match and if not fail the test
func (t *tester) check(values ...any) {
	if t.failed {
		return
	}

	if (len(values) & 1) == 1 {
		t.fail("There is an odd number (%d) of values", len(values))
		return
	}

	for i := 0; i < len(values); {
		a := values[i]
		b := values[i+1]
		i += 2

		if !reflect.DeepEqual(a, b) {
			t.fail("%v != %v", a, b)
			break
		}
	}
}

func (t *tester) fail(format string, data ...any) {
	_, _, line, ok := runtime.Caller(2)

	if ok {
		format = fmt.Sprintf("(line %d) %s", line, format)
	}

	t.Errorf(format, data...)
	t.failed = true
}

func (t *tester) shouldPanic() {
	if data := recover(); data == nil {
		t.fail("Test should panic but doesn't")
	}
}

func TestCap(t *testing.T) {
	data := make([]int, 0, 0)
	tt, stack := newTester[int](t)

	for i := 0; i < 10; i++ {
		data = append(data, i)
		stack.Push(i)
		tt.check(stack.Cap(), cap(data))
	}
}

func TestClear(t *testing.T) {
	tt, stack := newTester[int](t)
	stack.PushItems(1, 2, 3)
	tt.check(stack.Len(), 3)

	originalCap := stack.Cap()
	stack.Clear()

	tt.check(stack.Len(), 0, stack.Cap(), originalCap)
}

func TestClearWithCap(t *testing.T) {
	tt, stack := newTester[int](t)
	stack.PushItems(1, 2, 3)
	tt.check(stack.Len(), 3)
	stack.ClearWithCap(100)
	tt.check(stack.Len(), 0, stack.Cap(), 100)
}

func TestEmpty(t *testing.T) {
	tt, stack := newTester[int](t)
	tt.check(stack.Empty(), true)
	stack.Push(1)
	tt.check(stack.Empty(), false)
}

func TestFormat(t *testing.T) {
	tt, stack := newTester[float64](t)
	format := gs.NewFormat[float64]()

	format.Conversion = func(item *float64) string {
		n := *item * 100
		rounded := n + math.Copysign(0.5, n)
		val := math.Ceil(rounded) / 100
		return fmt.Sprintf("%.2f", val)
	}
	format.End = ")"
	format.Sep = " | "
	format.Start = "("

	stack.PushItems(.125, -.678)
	tt.check(stack.Format(format), "(-0.68 | 0.13)")

	format.TopFirst = false
	tt.check(stack.Format(format), "(0.13 | -0.68)")

	stack.Pop()
	tt.check(stack.Format(format), "(0.13)")
	stack.Clear()
	tt.check(stack.Format(format), "()")
}

func TestHasItems(t *testing.T) {
	tt, stack := newTester[int](t)
	tt.check(stack.HasItems(), false)
	stack.Push(1)
	tt.check(stack.HasItems(), true)
}

func TestLen(t *testing.T) {
	tt, stack := newTester[int](t)
	tt.check(stack.Len(), 0)
	stack.PushItems(1, 2, 3, 4, 5)
	tt.check(stack.Len(), 5)
	stack.Pop()
	tt.check(stack.Len(), 4)
}

func TestPop(t *testing.T) {
	tt, stack := newTester[int](t)
	defer tt.shouldPanic()

	stack.Push(1)
	tt.check(stack.Len(), 1)
	stack.Pop()
	tt.check(stack.Len(), 0)
	stack.Pop()
}

func TestPopAndReturn(t *testing.T) {
	tt, stack := newTester[int](t)
	defer tt.shouldPanic()

	stack.PushItems(1, 2, 3)

	a := stack.PopAndReturn()
	b := stack.PopAndReturn()
	c := stack.PopAndReturn()

	tt.check(a, 3, b, 2, c, 1)
	stack.PopAndReturn()
}

func TestPush(t *testing.T) {
	tt, stack := newTester[int](t)
	tt.check(stack.Len(), 0)
	stack.Push(100)
	tt.check(stack.Len(), 1, stack.Top(), 100)
}

func TestPushItems(t *testing.T) {
	tt, stack := newTester[int](t)
	tt.check(stack.Len(), 0)
	stack.PushItems(-1, -2, -3)
	tt.check(stack.Len(), 3, stack.Top(), -3)
}

func TestPushSlice(t *testing.T) {
	tt, stack := newTester[int](t)
	tt.check(stack.Len(), 0)
	stack.PushSlice([]int{-1, -2, -3})
	tt.check(stack.Len(), 3, stack.Top(), -3)
}

func TestString(t *testing.T) {
	tt, stack := newTester[int](t)
	stack.PushItems(10, 20)
	tt.check(fmt.Sprintf("%s", stack), "[20 10]")
	stack.Pop()
	tt.check(fmt.Sprintf("%s", stack), "[10]")
	stack.Clear()
	tt.check(fmt.Sprintf("%s", stack), "[]")
}

func TestTop(t *testing.T) {
	tt, stack := newTester[int](t)
	stack.Push(-1)
	tt.check(stack.Top(), -1)
	stack.Push(-2)
	tt.check(stack.Top(), -2)
	stack.Push(-3)
	tt.check(stack.Top(), -3)
}

func TestTopPointer(t *testing.T) {
	tt, stack := newTester[example](t)
	stack.PushItems(example{false, 10}, example{false, 20}, example{false, 30})

	topCopy := stack.Top()
	stack.TopPointer().flag = true
	stack.TopPointer().val = 40

	tt.check(topCopy.flag, false, stack.Top().flag, true)
	tt.check(topCopy.val, 30, stack.Top().val, 40)
}

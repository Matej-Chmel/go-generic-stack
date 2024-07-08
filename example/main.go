package main

import (
	"fmt"

	gs "github.com/Matej-Chmel/go-generic-stack"
)

func main() {
	stack := gs.New[float32]()
	stack.PushItems(.15, 1.5, 3)

	fmt.Println(stack.Top()) // 3
	*stack.TopPointer() = 3.14159

	fmt.Println(stack.PopAndReturn()) // 3.14159
	stack.Pop()

	fmt.Println(stack.Empty())    // false
	fmt.Println(stack.HasItems()) // true
	fmt.Println(stack.Len())      // 1
	fmt.Println(stack.Cap())      // 4

	stack.Push(10.56)
	stack.Push(20.99)

	fmt.Println(stack) // [20.99 10.56 0.15]

	format := gs.NewFormat[float32]()
	format.Conversion = func(item *float32) string {
		return fmt.Sprintf("%.1f", *item)
	}
	format.Start = "("
	format.Sep = ", "
	format.End = ")"

	fmt.Println(stack.Format(format)) // (21.0, 10.6, 0.2)

	stack.Clear()
	fmt.Println(stack) // []
}

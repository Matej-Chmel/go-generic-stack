# Stack
Simple stack data structure.

```go
package main

import (
    "fmt"

    gostack "github.com/Matej-Chmel/go-stack"
)

func main() {
    stack := gostack.Stack[int32]{}

    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    a := stack.Top()
    b, err := stack.Pop()

    if err == nil {
        fmt.Printf("%d == %d\n", a, b)
    }

    stack.Clear()

    if stack.Empty() {
        fmt.Println("Stack is empty")
    }
}
```
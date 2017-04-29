package collection

import (
	"testing"
	"fmt"
)

func TestStringStack(t *testing.T){
	var s Stack
	s.Push("world")
	s.Push("hello, ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
	// Output: hello, world
}
package main

import (
	"earthlyfisher.com/wyp/collection"
	"earthlyfisher.com/wyp/hellolib"
	"fmt"
)

func main() {
	fmt.Print("你好，世界！\n")
	fmt.Printf("2 和 3中最大的是 %d！", hellolib.Max(2, 3))
	fmt.Println("")

	text := "hello world!"
	for i, b := range []byte(text) {
		fmt.Println(i, ":", b)
	}

	//call stack example
	stackExample()
}

func stackExample() {
	var s collection.Stack
	s.Push("world")
	s.Push("hello, ")
	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
	fmt.Println()
}

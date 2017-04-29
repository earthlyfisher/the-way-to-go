// Package collection implements a generic stack.
package collection

// The zero value for Stack is an empty stack ready to use.
type Stack struct {
	data []interface{}
}

//interface{}在GO中是一个特殊的内建类型，类似于C/C++中的void*，但是包含了类型信息。
//所以你可以把任意的数据转换到interface{}，然后通过type assert从interface{}获取原有的数据。
//但是正如你所见，interface{}没有方法，那么也就是说，它不需要iface中的itab，因为不需要方法绑定。
//针对此，做了特殊修改，iface中的tab字段类型由itab指针变为了对应的具体实现类型的类型元数据指针。
// Push adds x to the top of the stack.
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// It’s a run-time error to call Pop on an empty stack.
func (s *Stack) Pop() interface{} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}

type StringStack struct {
	Stack
}

func (s *StringStack) Push(n string) {
	s.Stack.Push(n)
}

func (s *StringStack) Pop() string {
	return s.Stack.Pop().(string)
}

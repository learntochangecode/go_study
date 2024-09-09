package main

import "fmt"

// 支持泛型类型
// List 表示一个可以保存任何类型的值的单链表。
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Add(val)
	}
}

func (l *List[T]) Print() {
	fmt.Print(l.val, " ")
	if l.next != nil {
		l.next.Print()
	}
}

func main() {
	// 创建一个整数链表
	intList := &List[int]{val: 1}
	intList.Add(2)
	intList.Add(3)
	intList.Add(4)
	intList.Add(5)

	fmt.Println("Integer List:")
	intList.Print()
	fmt.Println()

	// 创建一个字符串链表
	strList := &List[string]{val: "Hello"}
	strList.Add("World")
	strList.Add("Go")
	strList.Add("Practice")

	fmt.Println("\nString List:")
	strList.Print()
	fmt.Println()
}

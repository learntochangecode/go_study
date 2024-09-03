package main

import "fmt"

type INTER interface {
	M()
}

func main() {
	var i INTER
	describe_inter(i)
	i.M()
}

// 运行时会报错 panic: runtime error: invalid memory address or nil pointer dereference
// 这是因为nil接口不保存值也不保存具体类型，nil接口调用方法产生运行时错误
func describe_inter(i INTER) {
	fmt.Printf("(%v, %T)\n", i, i)
}

package main

import "fmt"

func main() {
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)
}

// 空接口可保存任意类型的值，被用来处理未知类型的值
func describe1(i interface{}) {
	fmt.Printf("%v,%T\n", i, i)
}

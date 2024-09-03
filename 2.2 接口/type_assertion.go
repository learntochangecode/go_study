package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
	// 为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// 只保存一个值的情况下，若 i 并未保存 T 类型的值，该语句就会触发一个 panic。
	f = i.(float64) // panic
	fmt.Println(f)
}

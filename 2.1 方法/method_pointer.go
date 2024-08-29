package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 如果不用指针，那么Scale方法接收的是一个拷贝，只修改拷贝的值
// 如果需要修改原来的值，那么Scale方法接收指针类型
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// Scale方法的接收者是指针类型，所以调用时需要传递指针，此处是go的自动引用
	// 显式使用指针如下
	// vPointer := &v
	// vPointer.Scale(10)
	v.Scale(10)
	fmt.Println(v.Abs())
}

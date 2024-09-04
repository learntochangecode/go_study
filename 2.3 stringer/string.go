package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 相当于Java中的toString方法 实现之后通过Println方法可以自定义打印参数值
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

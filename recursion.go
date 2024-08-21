package main

import "fmt"

func main() {
	i := recursion(15)
	fmt.Println(i)
}
func recursion(n int64) int64 {
	if n == 0 {
		return 1
	} else if n > 0 {
		return n * recursion(n-1)
	} else {
		//抛出异常
		panic("n should be greater than 0")
	}
}

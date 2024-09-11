package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	// 当我们求和时，我们只关心总和，不关心x,y具体接收了哪一部分的和，只要保证x.y接收的值加起来等于总和即可，
	// 因此可以使用同一个信道进行传送。
	fmt.Println(x, y, x+y)
}

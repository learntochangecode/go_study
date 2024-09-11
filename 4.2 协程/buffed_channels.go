package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 当信道的缓冲区填满后，向其发送数据时会阻塞。
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 当缓冲区为空时，接受方会阻塞。
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}

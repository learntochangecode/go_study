package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("线程1在前..")
	say("线程2在前")
}

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}

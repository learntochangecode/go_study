package main

import "fmt"

func main() {
	map1 := make(map[string]string, 10)
	map2 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(map1, map2)
	// map无序，遍历时返回键值对顺序不固定
	for k, v := range map2 {
		fmt.Println(k, v)
	}
}

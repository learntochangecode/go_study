package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)
	split_string := strings.Fields(s)
	fmt.Println(split_string)
	for _, v := range split_string {
		word_count[v] = len(v)
	}
	return word_count
}

func main() {
	fmt.Println(WordCount("I am a boy"))
}

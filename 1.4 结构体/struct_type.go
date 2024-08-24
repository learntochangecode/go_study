package main

import "fmt"
import "time"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var (
		book1 Books
	)
	fmt.Printf("%+v\n", book1)
	book1.title = "西游记"
	book1.author = "吴承恩"
	book1.subject = "古典小说"
	book1.book_id = 1
	fmt.Printf("%+v", book1)
	time.Sleep(10 * time.Second)
}

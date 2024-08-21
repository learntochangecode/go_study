package main
import "fmt"
func main(){
	var i = 1
	var p *int = &i
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
}
package main
import "fmt"

func main(){
	var arr1 = [5] int {1,2,3,4,5}
	fmt.Println(arr1)
	arr2 := [2] string {"a","b"}
	fmt.Println(arr2)
	arr3 := [...] int {10,9,8,7,6,5,4,3,2,1}
	fmt.Println(arr3)
}
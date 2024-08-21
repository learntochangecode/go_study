package main
import "fmt"

func main(){
	const i int = 999
	const j = 998
	const k = 997
	const a,b,c = "aa",1,false
	fmt.Printf("%T\n",i)
	fmt.Printf("%T\n",j)
	fmt.Printf("%T,%T,%T",a,b,c)

}
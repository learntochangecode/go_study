package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// 这里必须用float64进行类型转换，如果不转换会用默认的ErrNegativeSqrt类型输出，但并没有实现String()方法
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // 返回错误
	}
	return math.Sqrt(x), nil // 返回平方根
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

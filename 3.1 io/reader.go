package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// io 包指定了 io.Reader 接口，它表示数据流的读取端。
	// Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。
	// io.Reader 接口有一个 Read 方法：
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

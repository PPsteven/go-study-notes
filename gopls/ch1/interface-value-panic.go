package main

import (
	"bytes"
	"fmt"
	"io"
)

// 当 debug 为false的时候，编译会通过，且 out != nil 也会是true，最终会导致panic
// 尽管一个nil的*bytes.Buffer指针有实现这个接口的方法，它也不满足这个接口具体的行为上的要求。
// 特别是这个调用违反了(*bytes.Buffer).Write方法的接收者非空的隐含先觉条件，所以将nil指针赋给这个接口是错误的。
// 接口值由接口的类型（动态类型 type）和这个类型具体的值（动态值 value）。
// 判断接口 是否等于 nil 的依据动态类型。
const debug = false

func main() {
	// 解决方法是 var buf io.Writer
	var buf *bytes.Buffer
	buf = nil
	fmt.Println(buf == nil)
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	// 当 out = (*bytes.Buffer)(nil) 的时候， out != nil 为true，但是由于(*bytes.Buffer).Write方法的接收者非空的隐含先觉条件，所以会报错。
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

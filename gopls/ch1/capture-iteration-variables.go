// 捕获迭代器变量需注意事项

package main

import "fmt"

var rmdirs []func()

func main() {
	for _, d := range []string{"a", "b", "desktop"} {
		rmdirs = append(rmdirs, func() {
			fmt.Printf("%v, %v \n", d, &d)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

// output:
//desktop
//desktop
//desktop

func main2() {
	for _, d := range []string{"a", "b", "desktop"} {
		d := d
		rmdirs = append(rmdirs, func() {
			fmt.Println(d)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

// output:
//a
//b
//desktop


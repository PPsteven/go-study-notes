// append 行为分析

package main

import "fmt"


func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func main() {
	var x []int
	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("%d addr=%p cap=%2d\t%v\n", i, x, cap(x), x)
	}
}

// append 行为分析
// 判断是否有足够空间保存新元素
// 1. 有：新元素赋值到末尾
// 2. 无：新建切片（cap 扩展为原来的2倍），并复制旧切片所有元素，最后新元素赋值到末尾
// 如下，可以看出append操作，可能返回不一致的底层数组引用
//0 cap=1	[0] addr 0xc0000b2008
//1 cap=2	[0 1] addr 0xc0000b2030
//2 cap=4	[0 1 2] addr 0xc0000b6020
//3 cap=4	[0 1 2 3] addr 0xc0000b6020
//4 cap=8	[0 1 2 3 4] addr 0xc0000b8040
//5 cap=8	[0 1 2 3 4 5] addr 0xc0000b8040
//6 cap=8	[0 1 2 3 4 5 6] addr 0xc0000b8040
//7 cap=8	[0 1 2 3 4 5 6 7] addr 0xc0000b8040
//8 cap=16	[0 1 2 3 4 5 6 7 8] addr 0xc0000ba000
//9 cap=16	[0 1 2 3 4 5 6 7 8 9] addr 0xc0000ba000

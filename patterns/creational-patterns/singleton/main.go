//Singleton creational design pattern restricts the instantiation of a type to a single object.

package main

import (
	"./singleton"
	"fmt"
)

// 初始化过程
func main()  {
	s := singleton.New()

	s["this"] = "that"

	s2 := singleton.New()

	fmt.Println("This is ", s2["this"])
}

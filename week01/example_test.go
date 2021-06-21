package unitTest

import (
	"fmt"
	"math/rand"
)


// Simple Example
func ExampleAdd(){
	fmt.Println(Add(1,2))
	fmt.Println(Add(2,2))
	// Output:
	// 3
	// 4
}


// Unorder output
func ExamplePerm() {
	for _, value := range rand.Perm(5) { // Example must use Perm method
		fmt.Println(value)
	}
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}
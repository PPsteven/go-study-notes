package main

import (
	"log"
	"time"
)

type FuncObject func()

func LogCostTime(fn FuncObject) FuncObject {
	return func() {
		n := time.Now()
		fn()
		log.Printf("Execution is completed with the result, cost %v\n", time.Since(n))
	}
}

func main() {
	f := LogCostTime(func() {
		time.Sleep(time.Second * 3)
	})

	f()
}

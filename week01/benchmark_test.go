package unitTest

import (
	"testing"
	"time"
)
// Simple Example
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(1, 2)
	}
}

func sleep(t int) {
		time.Sleep(time.Second * time.Duration(t))
}

// timer control
// go test -run=none -v -bench BenchmarkAdd2 benchmark_test.go unit_test.g
func BenchmarkAdd2(b *testing.B) {
	b.ResetTimer() // reset timer
	sleep(2)

	var n int
	for i := 0; i < 5; i++ {
		b.StopTimer() // stop timer
		sleep(1)
		b.StartTimer() // start timer
		n++
	}
}
// memory analyze
func heap() []byte {
	return make([]byte, 1024)
}

func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = heap()
	}
}

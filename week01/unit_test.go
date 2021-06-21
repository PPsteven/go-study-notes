package unitTest

import (
	"fmt"
	"testing"
	"time"
)

func Add(x, y int) int {
	return x + y
}

// Simple Testing
func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Log("log: add function success") // 仅失败 或 -v 输出
		t.FailNow()
		//t.Error() // Log -> Fail
		//t.Fatal() // Log -> FailNow
	}
}

// understanding t.Parallel
// TestA and TestB will run the same time
// go test --parallel 2
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

// table driven test
func Add2(x, y int) int {
	return x + y
}

func TestAdd2(t *testing.T) {
	fmt.Printf("start test TestAdd2")
	var test = []struct{
		x, y, expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{5, 3, 8},
	}

	for _, tt := range test {
		if actual := Add(tt.x, tt.y); actual != tt.expect {
			t.Errorf("add(%d, %d), expect: %d, actual:%d", tt.x, tt.y, tt.expect, actual)
		}
	}
}

// skip long runing test
// go test -short
func TestAdd3(t *testing.T) {
	fmt.Printf("start test TestAdd3")
	if testing.Short() {
		t.Skip("Skip long runing test")
	}
	time.Sleep(100 * time.Second)
}

// sometime wo need to do some necessary thing before or after testing, like connect and disconnect database.
// or wo need to control which code should run in main thred.
//func TestMain(m *testing.M){
//	// setup
//	fmt.Println("before test")
//	code := m.Run() // call test routine func
//	// teardown
//	fmt.Println("after test")
//	os.Exit(code)
//}

// Control each test, benchmarks,
//func TestMain(m *testing.M) {
//	// 这里是有坑，testing包经过升级改动，原先的MainStart就是现在的Main
//	// testing.Main 使用的特点是，可以有setup初始化，但是teardown不能自主控制
//	matchString := func(pat, str string) (bool, error) {
//		return true, nil
//	}
//
//	tests := []testing.InternalTest{
//		{"TestAdd", TestAdd},
//		{"TestAdd2", TestAdd2},
//	}
//
//	benchmarks := []testing.InternalBenchmark{}
//	examples := []testing.InternalExample{}
//
//	testing.Main(matchString, tests, benchmarks, examples)
//}


package go_reflect

import (
	"reflect"
	"testing"
)

func TestKindAndType(t *testing.T) {
	// Type 代码 Go 语言的类型对象：
	// Kind 是Type 的具体类型

	type Color string
	const (
		Red Color = "Red"
		Blue Color = "Blue"
		Yellow Color = "Yellow"
	)

	typeOfRead := reflect.TypeOf(Red)
	t.Logf("type of Red: Kind(%s), Type(%s)", typeOfRead.Kind(), typeOfRead.Name())
	// type of Red: Kind(string), Type(Color)

	// Type
	type PersonInfo struct {
		Name string `myjson:"name"`
		Sex  int64 `myjson:"sex"`
	}

	type StudentInfo struct {
		PersonInfo
		StuID int64
	}

	var stu StudentInfo
	typeOfStu := reflect.TypeOf(stu)
	t.Logf("Number of field: %d", typeOfStu.NumField())
	t.Logf("Field(i int) -> StructField: %#v", typeOfStu.Field(0))
	structField, _ := typeOfStu.FieldByName("PersonInfo")
	t.Logf("FieldByName(name string) -> StructField: %#v", structField)
	structField, _ = typeOfStu.FieldByNameFunc(func(s string) bool {
		if s == "Sex" { return true}
		return false
	})
	t.Logf("FieldByNameFunc( match func(string) bool) --> StructField: %#v", structField)
	t.Logf("FieldByIndex([]int) --> StructField: %#v", typeOfStu.FieldByIndex([]int{0, 1}))
	// Tag
	t.Logf("Get Tag By `myjson`: %s", structField.Tag.Get("myjson"))
}

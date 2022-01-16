package go_reflect

import (
	"reflect"
	"testing"
)

func TestGetValueOfSturct(t *testing.T){
	type Sex int
	const (
		Male Sex = iota
		Female
		)
	type PersonInfo struct {
		Name string `myjson:"name"`
		Sex  Sex `myjson:"sex"`
	}

	type StudentInfo struct {
		PersonInfo
		StuID int64
		Next *StudentInfo
	}

	stu := &StudentInfo{
		PersonInfo{"Tom", Male},
		302831,
		nil,
	}
	stu.Next = stu

	// 先传指针，后取实例，是为了能够后续修改
	valueOfStu := reflect.ValueOf(stu)
	valueOfStu = valueOfStu.Elem()
	t.Logf("Number of Value: %d", valueOfStu.NumField())
	t.Logf("Field(i) --> Value: %#v", valueOfStu.Field(0))
	t.Logf("FieldByName(i) --> Value: %#v", valueOfStu.FieldByName("StuID"))
	t.Logf("FieldByIndex([]int) --> Value: %#v", valueOfStu.FieldByIndex([]int{2,2,2,2,2,2,2,2,2,2,2,2}))


	// Get Value
	p := valueOfStu.Field(0)
	t.Logf("PersonInfo: %#v", p.Interface().(PersonInfo))
	t.Logf("PersonInfo.Name: %#v", p.Field(0).String())
	t.Logf("PersonInfo.Name: %#v", p.Field(0).Interface().(string))
	t.Logf("PersonInfo.Sex: %#v", p.FieldByName("Sex").Int())

	// Set Value
	valueOfStu.FieldByName("StuID").SetInt(3002832)
	t.Logf("%#v", stu)
}
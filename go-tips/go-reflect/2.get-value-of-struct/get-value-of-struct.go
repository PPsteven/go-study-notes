package main

import (
	"fmt"
	"reflect"
	"strings"
)

type CLC string

const (
	CLC_TYPE_A CLC = "马克思主义、列宁主义、毛泽东思想、邓小平理论"
	CLC_TYPE_B CLC = "哲学、宗教"
	CLC_TYPE_C CLC = "社会科学总论"
)

type Goods struct {
	Name  string
	Price int64
}

type Book struct {
	Goods
	Author string `cn:"作者" orm:"author"`
	CLC    CLC    `cn:"中图分类号" orm:"clc"`
}

func GetValueOfSturct() {
	book := Book{
		Goods:  Goods{"悲惨世界", 300},
		Author: "雨果",
		CLC:    CLC_TYPE_A,
	}

	valueOfBook := reflect.ValueOf(book)

	n := valueOfBook.NumField()
	fmt.Printf("Number of Value: %d\n", n)
	// Number of Value: 3

	for i := 0; i < n; i++ {
		fmt.Printf("field[%d] --> %v\n", i, valueOfBook.Field(i))
	}

	fmt.Println(strings.Repeat("=", 20))

	names := []string{"Goods", "Author", "CLC"}
	for _, name := range names {
		fmt.Printf("field name(%s) --> %v\n", name, valueOfBook.FieldByName(name))
	}

	fmt.Println(strings.Repeat("=", 20))

	goods := valueOfBook.FieldByNameFunc(func(s string) bool {
		if strings.Contains(s, "Good") {
			return true
		}
		return false
	})
	fmt.Printf("fuzzy[Good] --> %v\n", goods)

	fmt.Println(strings.Repeat("=", 20))

	// Get Value
	// turn Value to go type
	goods = valueOfBook.Field(0)
	fmt.Printf("Goods: %v\n", goods.Interface().(Goods))
	fmt.Printf("Goods.Name: %s\n", goods.Field(0).String())
	fmt.Printf("Goods.Name: %s\n", goods.Field(0).Interface().(string)) // another way
	fmt.Printf("Goods.Price: %d\n", goods.FieldByName("Price").Int())
	//Goods: {悲惨世界 300}
	//Goods.Name: 悲惨世界
	//Goods.Name: 悲惨世界
	//Goods.Price: 300

	fmt.Println(strings.Repeat("=", 20))

	// Set Value
	bookPtr := &Book{
		Goods:  Goods{"悲惨世界", 300},
		Author: "雨果",
		CLC:    CLC_TYPE_A,
	}
	// 先传指针，后取实例，是为了能够修改
	valueOfBook = reflect.ValueOf(bookPtr)
	valueOfBook = valueOfBook.Elem()

	goods = valueOfBook.FieldByName("Goods")
	fmt.Printf("before: %v\n", *bookPtr)
	goods.FieldByName("Price").SetInt(40)
	fmt.Printf("after: %v\n", *bookPtr)
}

func ABC() {
	var values []interface{}
	var stmt string
	book := Book{
		Author: "雨果",
		CLC:    CLC_TYPE_A,
	}

	t := reflect.TypeOf(book)
	v := reflect.ValueOf(book)
	for i:= 0; i< v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("orm")
		if tag != "" {
			values = append(values, v.Field(i).Interface())
			stmt = stmt + fmt.Sprintf("%s,", tag)
		}
	}
	fmt.Println("stmt: ", stmt)
	fmt.Println("values: ", values)
}
func main() {
	//GetValueOfSturct()
	ABC()
}

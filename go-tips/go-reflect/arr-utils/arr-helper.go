package arr_utils

import (
	"reflect"
)

func pop(arr interface{}, i int) {
	v := reflect.ValueOf(arr).Elem()
	v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
}

func insert(arr interface{}, i int, val interface{}) {
	v := reflect.ValueOf(arr).Elem()
	v.Set(reflect.AppendSlice(v.Slice(0, i+1), v.Slice(i, v.Len())))
	v.Index(i).Set(reflect.ValueOf(val))
}

func deleteCopy(slice interface{}, i int) {
	v := reflect.ValueOf(slice).Elem()
	tmp := reflect.MakeSlice(v.Type(), 0, v.Len()-1)
	v.Set(
		reflect.AppendSlice(
			reflect.AppendSlice(tmp, v.Slice(0, i)),
			v.Slice(i+1, v.Len())))
}

func insertCopy(slice interface{}, i int, val interface{}) {
	v := reflect.ValueOf(slice).Elem()
	tmp := reflect.MakeSlice(v.Type(), 0, v.Len()+1)
	v.Set(reflect.AppendSlice(
		reflect.AppendSlice(tmp, v.Slice(0, i+1)),
		v.Slice(i, v.Len())))
	v.Index(i).Set(reflect.ValueOf(val))
}
//
//func main() {
//	arr := []int{0, 1, 2, 3, 4, 5, 6}[:6]
//	brr := arr
//
//	fmt.Println("arr:", arr, "brr:", brr)
//	insert(&arr, 2, 8)
//	fmt.Println("arr:", arr, "brr:", brr)
//	Delete(&arr, 5)
//	fmt.Println("arr:", arr, "brr:", brr)
//
//	fmt.Println("\nCopy Version\n")
//
//	arr = []int{0, 1, 2, 3, 4, 5, 6}[:6]
//	brr = arr
//
//	fmt.Println("arr:", arr, "brr:", brr)
//	insert_copy(&arr, 2, 8)
//	fmt.Println("arr:", arr, "brr:", brr)
//	fmt.Println("brr was unchanged, setting to arr")
//	brr = arr
//	delete_copy(&arr, 5)
//	fmt.Println("arr:", arr, "brr:", brr)
//
//}

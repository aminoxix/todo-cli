package main

import (
	"fmt"

	// "reflect"
	"strconv"
)

func main() {
	// models are different
	// fmt.Println(strconv.FormatBool(compare([]int{1, 2, 3}, []int{1, 2, 3})))

 	// models are same
	fmt.Println(strconv.FormatBool(compare(1, 1)))

}


// func compare(first, second interface{}) bool {
// 	// if reflect.TypeOf(first) == reflect.TypeOf(second) {
// 	if fmt.Sprintf("%T", first) == fmt.Sprintf("%T", second) {
// 		fmt.Println("Same type")
// 	} else {
// 		fmt.Println("Oops! Different type")
// 	}
// 	return first == second
// }

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string | ~bool | ~complex64 | ~complex128 
}
func compare[T Ordered](first, second T) bool {
	return first == second
}
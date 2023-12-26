package generichelper

import (
	"fmt"
)

func testNoType_any[T any]() string {
	if TypeOf[T]() == TypeOfNoType {
		return "NoType"
	}
	return TypeOf[T]().Kind().String()
}

func testNoType_comparable[T comparable]() string {
	if TypeOf[T]() == TypeOfNoType {
		return "NoType"
	}
	return TypeOf[T]().Kind().String()
}

func ExampleNoType() {
	fmt.Println(testNoType_any[NoType]())
	fmt.Println(testNoType_any[struct{}]())
	fmt.Println(testNoType_comparable[NoType]())
	fmt.Println(testNoType_comparable[int]())
	// Output:
	// NoType
	// struct
	// NoType
	// int
}

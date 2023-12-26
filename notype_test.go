package generichelper

import (
	"fmt"
)

func testNoType[T any]() string {
	if TypeOf[T]() == TypeOfNoType {
		return "NoType"
	}
	return TypeOf[T]().Kind().String()
}

func ExampleNoType() {
	fmt.Println(testNoType[NoType]())
	fmt.Println(testNoType[int]())
	// Output:
	// NoType
	// int
}

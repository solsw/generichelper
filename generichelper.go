package generichelper

import (
	"reflect"
)

// ZeroValue returns T's [zero value].
//
// [zero value]: https://go.dev/ref/spec#The_zero_value
func ZeroValue[T any]() T {
	var t0 T
	return t0
}

// DeepEqual is a generic wrapper around [reflect.DeepEqual].
func DeepEqual[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}

// Ternary mimics [ternary conditional operation].
//
// [ternary conditional operation]: https://en.wikipedia.org/wiki/Ternary_conditional_operator
func Ternary[T any](condition bool, trueT, falseT T) T {
	// https://golang.org/doc/faq#Does_Go_have_a_ternary_form
	if condition {
		return trueT
	}
	return falseT
}

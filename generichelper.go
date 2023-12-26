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

// TypeOf returns T's [reflect.Type].
func TypeOf[T any]() reflect.Type {
	var t0 T
	return reflect.TypeOf(t0)
}

// DeepEqual is a generic wrapper around [reflect.DeepEqual].
func DeepEqual[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}

// Is determines whether 'x' [is] of type T.
//
// [is]: https://go.dev/ref/spec#Type_assertions
func Is[T any](x any) bool {
	_, ok := x.(T)
	return ok
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

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

// IsZeroValue reports whether t equals T's [zero value] using [reflect.DeepEqual].
//
// [zero value]: https://go.dev/ref/spec#The_zero_value
func IsZeroValue[T any](t T) bool {
	return DeepEqual(t, ZeroValue[T]())
}

// Is determines whether 'x' [is] of type T.
//
// Since Is relies on a [type assertion], a nil 'x' carries no dynamic type
// and yields 'false' for any T, including T == any.
//
// [is]: https://go.dev/ref/spec#Type_assertions
// [type assertion]: https://go.dev/ref/spec#Type_assertions
func Is[T any](x any) bool {
	_, ok := x.(T)
	return ok
}

// Ternary mimics [ternary conditional operation].
//
// Unlike a real ternary operator, Ternary does not short-circuit:
// both trueT and falseT are evaluated by the caller before Ternary is called.
// Do not use Ternary for operands that have side effects
// or that are unsafe to evaluate (e.g. dereferencing a possibly-nil pointer).
//
// [ternary conditional operation]: https://en.wikipedia.org/wiki/Ternary_conditional_operator
func Ternary[T any](condition bool, trueT, falseT T) T {
	// https://golang.org/doc/faq#Does_Go_have_a_ternary_form
	if condition {
		return trueT
	}
	return falseT
}

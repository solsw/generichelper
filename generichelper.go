// Package generichelper contains helpers for Go generics.
package generichelper

import (
	"reflect"
)

// ZeroValue returns T's zero value (https://go.dev/ref/spec#The_zero_value).
func ZeroValue[T any]() T {
	var t0 T
	return t0
}

// DeepEqual is a generic wrapper for reflect.DeepEqual.
func DeepEqual[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}

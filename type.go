package generichelper

import (
	"reflect"
)

// SameType determines whether T's and O's types are the same.
func SameType[T, O any]() bool {
	return reflect.TypeFor[T]() == reflect.TypeFor[O]()
}

// NoType is a sentinel type to denote that this [type parameter]
// needs not to be processed (see ExampleNoType).
// NoType satisfies [type constraints] such as any, comparable, and ~struct{}.
//
// [type constraints]: https://go.dev/ref/spec#Type_constraints
// [type parameter]: https://go.dev/ref/spec#Type_parameter_declarations
type NoType struct{}

// IsNoType determines whether T's type is [NoType].
func IsNoType[T any]() bool {
	return SameType[T, NoType]()
}

// TypeHoldsType reports whether T and O are the same type or O is an
// interface that T [implements].
//
// [implements]: https://go.dev/ref/spec#Implementing_an_interface
func TypeHoldsType[T, O any]() bool {
	tType := reflect.TypeFor[T]()
	oType := reflect.TypeFor[O]()
	if oType.Kind() == reflect.Interface {
		return tType.Implements(oType)
	}
	return tType == oType
}

// TypeMeetsType returns 'true' if:
//   - [TypeHoldsType] returns 'true' or
//   - type T can be [converted] to type O.
//
// Otherwise, TypeMeetsType returns 'false'.
//
// [converted]: https://go.dev/ref/spec#Conversions
func TypeMeetsType[T, O any]() bool {
	if TypeHoldsType[T, O]() {
		return true
	}
	// Type-level convertibility: unlike reflect.Value.CanConvert, this does
	// not depend on a particular value (e.g. a slice's length for slice-to-array
	// conversions), matching the documented type-based contract.
	return reflect.TypeFor[T]().ConvertibleTo(reflect.TypeFor[O]())
}

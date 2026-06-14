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
// NoType may be used to [instantiate] [type parameter]
// with 'any' or 'comparable' [type constraint] only.
//
// [instantiate]: https://go.dev/ref/spec#Instantiations
// [type constraint]: https://go.dev/ref/spec#Type_constraints
// [type parameter]: https://go.dev/ref/spec#Type_parameter_declarations
type NoType struct{}

// IsNoType determines whether T's type is [NoType].
func IsNoType[T any]() bool {
	return SameType[T, NoType]()
}

// TypeHoldsType reports whether type T holds other type O.
//
// TypeHoldsType returns 'true' if:
//   - [type assertion] T.(O) holds or
//   - type T [implements] interface type O.
//
// Otherwise, TypeHoldsType returns 'false'.
//
// [implements]: https://go.dev/ref/spec#Implementing_an_interface
// [type assertion]: https://go.dev/ref/spec#Type_assertions
func TypeHoldsType[T, O any]() bool {
	if reflect.TypeFor[T]().Kind() != reflect.Interface {
		// ZeroValue[T]() carries T's dynamic type even when nil,
		// so the assertion reports whether T holds O.
		_, isO := any(ZeroValue[T]()).(O)
		return isO
	}
	// T is an interface; its zero value is an untyped nil that cannot be
	// asserted, so decide by checking whether T implements interface O.
	oType := reflect.TypeFor[O]()
	if oType.Kind() != reflect.Interface {
		return false
	}
	return reflect.TypeFor[T]().Implements(oType)
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

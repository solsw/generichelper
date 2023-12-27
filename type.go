package generichelper

import (
	"reflect"
)

// NoType is a sentinel type to denote that this [type parameter]
// needs not to be processed (see ExampleNoType).
// NoType may be used to [instantiate] [type parameter]
// with 'any' or 'comparable' [type constraint] only.
//
// [instantiate]: https://go.dev/ref/spec#Instantiations
// [type constraint]: https://go.dev/ref/spec#Type_constraints
// [type parameter]: https://go.dev/ref/spec#Type_parameter_declarations
type NoType struct{}

// TypeOf returns T's [reflect.Type].
func TypeOf[T any]() reflect.Type {
	var t0 T
	return reflect.TypeOf(t0)
}

var (
	typeOfNoType reflect.Type = TypeOf[NoType]()
	typeOfString reflect.Type = TypeOf[string]()
)

// IsNoType determines whether T's type is [NoType].
func IsNoType[T any]() bool {
	return TypeOf[T]() == typeOfNoType
}

// IsString determines whether T's type is [string].
func IsString[T any]() bool {
	return TypeOf[T]() == typeOfString
}

func typeHoldsTypePrim[T, O any]() (isO bool, t0 any, oType reflect.Type) {
	t0 = ZeroValue[T]()
	if t0 != nil {
		// T is a value type
		_, isO = t0.(O)
		return isO, t0, nil
	}
	// t0 is nil here for interfaces
	oType = reflect.TypeOf((*O)(nil)).Elem()
	if oType.Kind() == reflect.Interface {
		tType := reflect.TypeOf((*T)(nil)).Elem()
		return tType.Implements(oType), t0, oType
	}
	return false, t0, oType
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
	isO, _, _ := typeHoldsTypePrim[T, O]()
	return isO
}

// TypeIsType reports whether type T is other type O.
//
// TypeIsType returns 'true' if:
//   - [TypeHoldsType] returns 'true' or
//   - type T can be [converted] to type O.
//
// Otherwise, TypeIsType returns 'false'.
//
// [converted]: https://go.dev/ref/spec#Conversions
func TypeIsType[T, O any]() bool {
	holdsO, t0, oType := typeHoldsTypePrim[T, O]()
	if holdsO {
		return true
	}
	if t0 == nil {
		return false
	}
	if oType == nil {
		oType = reflect.TypeOf((*O)(nil)).Elem()
	}
	tValue := reflect.ValueOf(t0)
	return tValue.CanConvert(oType)
}
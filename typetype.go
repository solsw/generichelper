package generichelper

import (
	"reflect"
)

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

// TypeHoldsType checks whether type T holds other type O.
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

// TypeIsType checks whether type T is other type O.
//
// TypeIsType returns 'true' if:
//   - [TypeHoldsType] returns 'true' or
//   - type T can be [converted] to type O.
//
// Otherwise, TypeIsType returns 'false'.
//
// [converted]: https://go.dev/ref/spec#Conversions
func TypeIsType[T, O any]() bool {
	isO, t0, oType := typeHoldsTypePrim[T, O]()
	if isO {
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

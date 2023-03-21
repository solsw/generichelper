package generichelper

import (
	"reflect"
)

// TypeIsType checks whether type T is other type O.
//
// TypeIsType returns 'true' if:
//   - type T can be [converted] to type O;
//   - [type assertion] T.(O) holds;
//   - type T [implements] interface type O.
//
// Otherwise, TypeIsType returns 'false'.
//
// [converted]: https://go.dev/ref/spec#Conversions
// [type assertion]: https://go.dev/ref/spec#Type_assertions
// [implements]: https://go.dev/ref/spec#Interface_types
func TypeIsType[T, O any]() bool {
	var t0 any = ZeroValue[T]()
	if t0 == nil {
		// t0 is nil for pointers, functions, interfaces, slices, channels, and maps
		oType := reflect.TypeOf((*O)(nil)).Elem()
		if oType.Kind() == reflect.Interface {
			tType := reflect.TypeOf((*T)(nil)).Elem()
			return tType.Implements(oType)
		}
	} else {
		_, isO := t0.(O)
		if isO {
			return true
		}
		tValue := reflect.ValueOf(t0)
		oType := reflect.TypeOf((*O)(nil)).Elem()
		if tValue.CanConvert(oType) {
			return true
		}
	}
	return false
}

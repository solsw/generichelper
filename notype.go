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

// TypeOfNoType is [reflect.Type] of [NoType].
var TypeOfNoType reflect.Type = TypeOf[NoType]()

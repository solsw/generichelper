# generichelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/generichelper.svg)](https://pkg.go.dev/github.com/solsw/generichelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/generichelper)

Helpers for Go's [generics](https://go.dev/doc/tutorial/generics).

## Installation

```sh
go get github.com/solsw/generichelper
```

Requires Go 1.22 or later (the package uses [`reflect.TypeFor`](https://pkg.go.dev/reflect#TypeFor)).

```go
import "github.com/solsw/generichelper"
```

## Overview

`generichelper` collects small, dependency-free helpers that fill gaps in Go's
generics: obtaining zero values, type-safe equality, lightweight runtime/type
checks, a ternary expression, and a 2-tuple.

## API

### Values

#### `ZeroValue[T any]() T`

Returns `T`'s [zero value](https://go.dev/ref/spec#The_zero_value).

```go
n := generichelper.ZeroValue[int]()    // 0
s := generichelper.ZeroValue[string]() // ""
p := generichelper.ZeroValue[*int]()   // nil
```

#### `IsZeroValue[T any](t T) bool`

Reports whether `t` equals `T`'s zero value.

`IsZeroValue` is implemented with [`reflect.DeepEqual`](https://pkg.go.dev/reflect#DeepEqual),
so it allocates and reflects on every call and inherits `reflect.DeepEqual`'s
comparison semantics.

```go
generichelper.IsZeroValue(0)   // true
generichelper.IsZeroValue("")  // true
generichelper.IsZeroValue("x") // false
```

#### `DeepEqual[T any](x, y T) bool`

A generic, type-safe wrapper around `reflect.DeepEqual`: `x` and `y` are
constrained to the same type `T`.

```go
generichelper.DeepEqual([]int{1, 2}, []int{1, 2}) // true
```

### Runtime checks

#### `Is[T any](x any) bool`

Reports whether `x` [is](https://go.dev/ref/spec#Type_assertions) of type `T`.

Because `Is` relies on a type assertion, a `nil` `x` carries no dynamic type and
yields `false` for any `T`, including `T == any`.

```go
generichelper.Is[string]("hi") // true
generichelper.Is[string](42)   // false
generichelper.Is[any](nil)     // false
```

### Conditional

#### `Ternary[T any](condition bool, trueT, falseT T) T`

Mimics the [ternary conditional operator](https://en.wikipedia.org/wiki/Ternary_conditional_operator):
returns `trueT` when `condition` is `true`, otherwise `falseT`.

> **Note:** `Ternary` does not short-circuit. Both `trueT` and `falseT` are
> evaluated by the caller before `Ternary` is called, so do not use it for
> operands that have side effects or that are unsafe to evaluate (e.g.
> dereferencing a possibly-nil pointer).

```go
label := generichelper.Ternary(n%2 == 0, "even", "odd")
```

### Type relationships

#### `SameType[T, O any]() bool`

Reports whether `T` and `O` are the same type. Type aliases are the same type;
defined types with the same underlying type are not.

```go
generichelper.SameType[string, string]() // true
generichelper.SameType[int, string]()    // false
```

#### `TypeHoldsType[T, O any]() bool`

Reports whether type `T` holds type `O`, i.e. returns `true` when the
[type assertion](https://go.dev/ref/spec#Type_assertions) `T.(O)` holds or
when `T` [implements](https://go.dev/ref/spec#Implementing_an_interface) the
interface type `O`.

```go
generichelper.TypeHoldsType[int, int]()              // true
generichelper.TypeHoldsType[*os.File, io.Writer]()   // true (implements)
generichelper.TypeHoldsType[int, string]()           // false
```

#### `TypeMeetsType[T, O any]() bool`

Returns `true` when `TypeHoldsType[T, O]()` is `true` **or** when `T` can be
[converted](https://go.dev/ref/spec#Conversions) to `O`. Convertibility is
decided at the type level, so it does not depend on any particular value (for
example, a slice's length in a slice-to-array conversion).

```go
generichelper.TypeMeetsType[int, string]()   // true (int convertible to string)
generichelper.TypeMeetsType[[]int, [3]int]() // true (valid conversion)
generichelper.TypeMeetsType[bool, string]()  // false
```

#### `NoType` and `IsNoType[T any]() bool`

`NoType` is a sentinel type used to mark a type parameter that should not be
processed. It may instantiate type parameters constrained by `any` or
`comparable` only. `IsNoType[T]()` reports whether `T` is `NoType`.

```go
func process[T any](v T) {
    if generichelper.IsNoType[T]() {
        return // nothing to process
    }
    // ...
}
```

### Tuple

#### `Tuple2[T1, T2 any]` and `NewTuple2`

A 2-tuple (pair) with exported fields `Item1` and `Item2`. `NewTuple2` is a
convenience constructor whose arguments can be supplied directly by a function
returning two values.

```go
p := generichelper.NewTuple2(1, "one")
fmt.Println(p.Item1, p.Item2) // 1 one
```

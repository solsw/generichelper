package generichelper

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestTypeOf_int(t *testing.T) {
	var v0 int
	want := reflect.TypeOf(v0)
	if got := TypeOf[int](); got != want {
		t.Errorf("TypeOf[int]() = %v, want %v", got, want)
	}
}

func TestTypeOf_string(t *testing.T) {
	var v0 string
	want := reflect.TypeOf(v0)
	if got := TypeOf[string](); got != want {
		t.Errorf("TypeOf[string]() = %v, want %v", got, want)
	}
}

func ExampleSameType_string() {
	fmt.Println(SameType[string, NoType]())
	fmt.Println(SameType[string, string]())
	fmt.Println(SameType[int, string]())
	type s2 = string
	fmt.Println(SameType[string, s2]())
	type s3 string
	fmt.Println(SameType[s3, string]())
	// Output:
	// false
	// true
	// false
	// true
	// false
}

func testNoType_any[T any]() string {
	if IsNoType[T]() {
		return "NoType"
	}
	return TypeOf[T]().Kind().String()
}

func testNoType_comparable[T comparable]() string {
	if IsNoType[T]() {
		return "NoType"
	}
	return TypeOf[T]().Kind().String()
}

func ExampleNoType() {
	fmt.Println(testNoType_any[NoType]())
	fmt.Println(testNoType_any[struct{}]())
	fmt.Println(testNoType_comparable[NoType]())
	fmt.Println(testNoType_comparable[int]())
	// Output:
	// NoType
	// struct
	// NoType
	// int
}

func TestTypeHoldsType_int_int(t *testing.T) {
	want := true
	if got := TypeHoldsType[int, int](); got != want {
		t.Errorf("TypeHoldsType[int, int]() = %v, want %v", got, want)
	}
}

func TestTypeHoldsType_int_string(t *testing.T) {
	want := false
	if got := TypeHoldsType[int, string](); got != want {
		t.Errorf("TypeHoldsType[int, string]() = %v, want %v", got, want)
	}
}

func TestTypeHoldsType_any_int(t *testing.T) {
	want := false
	if got := TypeHoldsType[any, int](); got != want {
		t.Errorf("TypeHoldsType[any, int]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_int_int(t *testing.T) {
	want := true
	if got := TypeMeetsType[int, int](); got != want {
		t.Errorf("TypeMeetsType[int, int]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_bool_string(t *testing.T) {
	want := false
	if got := TypeMeetsType[bool, string](); got != want {
		t.Errorf("TypeMeetsType[bool, string]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_int_string(t *testing.T) {
	want := true
	if got := TypeMeetsType[int, string](); got != want {
		t.Errorf("TypeMeetsType[int, string]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_io_ReadWriter_io_Writer(t *testing.T) {
	want := true
	if got := TypeMeetsType[io.ReadWriter, io.Writer](); got != want {
		t.Errorf("TypeMeetsType[io.ReadWriter, io.Writer]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_io_Writer_io_ReadWriter(t *testing.T) {
	want := false
	if got := TypeMeetsType[io.Writer, io.ReadWriter](); got != want {
		t.Errorf("TypeMeetsType[io.Writer, io.ReadWriter]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_os_File_io_Writer(t *testing.T) {
	want := false
	if got := TypeMeetsType[os.File, io.Writer](); got != want {
		t.Errorf("TypeMeetsType[os.File, io.Writer]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_Ptr_os_File_io_Writer(t *testing.T) {
	want := true
	if got := TypeMeetsType[*os.File, io.Writer](); got != want {
		t.Errorf("TypeMeetsType[*os.File, io.Writer]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_uint_reflect_Kind(t *testing.T) {
	want := true
	if got := TypeMeetsType[uint, reflect.Kind](); got != want {
		t.Errorf("TypeMeetsType[uint, reflect.Kind]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_reflect_Kind_uint(t *testing.T) {
	want := true
	if got := TypeMeetsType[reflect.Kind, uint](); got != want {
		t.Errorf("TypeMeetsType[reflect.Kind, uint]() = %v, want %v", got, want)
	}
}

func TestTypeMeetsType_reflect_Kind_fmt_Stringer(t *testing.T) {
	want := true
	if got := TypeMeetsType[reflect.Kind, fmt.Stringer](); got != want {
		t.Errorf("TypeMeetsType[reflect.Kind, fmt.Stringer]() = %v, want %v", got, want)
	}
}

type SliceWithString []int

func (SliceWithString) String() string {
	return ""
}

func TestTypeMeetsType_SliceWithString_fmt_Stringer(t *testing.T) {
	want := true
	if got := TypeMeetsType[SliceWithString, fmt.Stringer](); got != want {
		t.Errorf("TypeMeetsType[SliceWithString, fmt.Stringer]() = %v, want %v", got, want)
	}
}

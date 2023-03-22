package generichelper

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

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

func TestTypeIsType_int_int(t *testing.T) {
	want := true
	if got := TypeIsType[int, int](); got != want {
		t.Errorf("TypeIsType[int, int]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_bool_string(t *testing.T) {
	want := false
	if got := TypeIsType[bool, string](); got != want {
		t.Errorf("TypeIsType[bool, string]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_int_string(t *testing.T) {
	want := true
	if got := TypeIsType[int, string](); got != want {
		t.Errorf("TypeIsType[int, string]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_io_ReadWriter_io_Writer(t *testing.T) {
	want := true
	if got := TypeIsType[io.ReadWriter, io.Writer](); got != want {
		t.Errorf("TypeIsType[io.ReadWriter, io.Writer]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_io_Writer_io_ReadWriter(t *testing.T) {
	want := false
	if got := TypeIsType[io.Writer, io.ReadWriter](); got != want {
		t.Errorf("TypeIsType[io.Writer, io.ReadWriter]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_os_File_io_Writer(t *testing.T) {
	want := false
	if got := TypeIsType[os.File, io.Writer](); got != want {
		t.Errorf("TypeIsType[os.File, io.Writer]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_Ptr_os_File_io_Writer(t *testing.T) {
	want := true
	if got := TypeIsType[*os.File, io.Writer](); got != want {
		t.Errorf("TypeIsType[*os.File, io.Writer]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_uint_reflect_Kind(t *testing.T) {
	want := true
	if got := TypeIsType[uint, reflect.Kind](); got != want {
		t.Errorf("TypeIsType[uint, reflect.Kind]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_reflect_Kind_uint(t *testing.T) {
	want := true
	if got := TypeIsType[reflect.Kind, uint](); got != want {
		t.Errorf("TypeIsType[reflect.Kind, uint]() = %v, want %v", got, want)
	}
}

func TestTypeIsType_reflect_Kind_fmt_Stringer(t *testing.T) {
	want := true
	if got := TypeIsType[reflect.Kind, fmt.Stringer](); got != want {
		t.Errorf("TypeIsType[reflect.Kind, fmt.Stringer]() = %v, want %v", got, want)
	}
}

type SliceWithString []int

func (SliceWithString) String() string {
	return ""
}

func TestTypeIsType_SliceWithString_fmt_Stringer(t *testing.T) {
	want := true
	if got := TypeIsType[SliceWithString, fmt.Stringer](); got != want {
		t.Errorf("TypeIsType[SliceWithString, fmt.Stringer]() = %v, want %v", got, want)
	}
}

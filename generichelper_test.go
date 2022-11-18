package generichelper

import (
	"reflect"
	"testing"
)

func TestZeroValue_bool(t *testing.T) {
	var want bool
	if got := ZeroValue[bool](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[bool]() = %v, want %v", got, want)
	}
}

func TestZeroValue_int(t *testing.T) {
	var want int
	if got := ZeroValue[int](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[int]() = %v, want %v", got, want)
	}
}

func TestZeroValue_string(t *testing.T) {
	var want string
	if got := ZeroValue[string](); !reflect.DeepEqual(got, want) {
		t.Errorf(`ZeroValue[string]() = "%v", want "%v"`, got, want)
	}
}

func TestZeroValue_array(t *testing.T) {
	var want [2]int
	if got := ZeroValue[[2]int](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[[2]int]() = %v, want %v", got, want)
	}
}

func TestZeroValue_slice(t *testing.T) {
	var want []int
	if got := ZeroValue[[]int](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[[]int]() = %v, want %v", got, want)
	}
}

func TestZeroValue_struct(t *testing.T) {
	var want struct{}
	if got := ZeroValue[struct{}](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[struct{}]() = %v, want %v", got, want)
	}
}

func TestZeroValue_pointer(t *testing.T) {
	var want *struct{}
	if got := ZeroValue[*struct{}](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[*struct{}]() = %v, want %v", got, want)
	}
}

func TestZeroValue_func(t *testing.T) {
	var want func(x int) int
	if got := ZeroValue[func(x int) int](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[func(x int) int]() = %p, want %p", got, want)
	}
}

func TestZeroValue_interface(t *testing.T) {
	var want any
	if got := ZeroValue[any](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[any]() = %v, want %v", got, want)
	}
}

func TestZeroValue_map(t *testing.T) {
	var want map[int]string
	if got := ZeroValue[map[int]string](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[map[int]string]() = %v, want %v", got, want)
	}
}

func TestZeroValue_chan(t *testing.T) {
	var want chan int
	if got := ZeroValue[chan int](); !reflect.DeepEqual(got, want) {
		t.Errorf("ZeroValue[chan int]() = %v, want %v", got, want)
	}
}

func TestDeepEqual_int(t *testing.T) {
	want := true
	if got := DeepEqual(2, 2); !reflect.DeepEqual(got, want) {
		t.Errorf("DeepEqual[int]() = %v, want %v", got, want)
	}
}

func TestDeepEqual_string(t *testing.T) {
	want := true
	if got := DeepEqual("2", "2"); !reflect.DeepEqual(got, want) {
		t.Errorf("DeepEqual[string]() = %v, want %v", got, want)
	}
}

package generichelper

import (
	"errors"
	"math"
	"reflect"
	"testing"
	"time"
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

func TestDeepEqual_float64(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal",
			args: args{
				x: 1,
				y: 1,
			},
			want: true,
		},
		{name: "not equal",
			args: args{
				x: 1,
				y: 2,
			},
			want: false,
		},
		{name: "NaNNaN",
			args: args{
				x: math.NaN(),
				y: math.NaN(),
			},
			want: false,
		},
		{name: "NaN1",
			args: args{
				x: math.NaN(),
				y: 1,
			},
			want: false,
		},
		{name: "1NaN",
			args: args{
				x: 1,
				y: math.NaN(),
			},
			want: false,
		},
		{name: "+Inf+Inf",
			args: args{
				x: math.Inf(+1),
				y: math.Inf(+1),
			},
			want: true,
		},
		{name: "-Inf-Inf",
			args: args{
				x: math.Inf(-1),
				y: math.Inf(-1),
			},
			want: true,
		},
		{name: "+InfNaN",
			args: args{
				x: math.Inf(+1),
				y: math.NaN(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepEqual(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("DeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs_any(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "int",
			args: args{x: 1},
			want: true,
		},
		{name: "string",
			args: args{x: "two"},
			want: true,
		},
		{name: "error",
			args: args{x: errors.New("")},
			want: true,
		},
		{name: "slice",
			args: args{x: []int{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is[any](tt.args.x); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs_string(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "int",
			args: args{x: 1},
			want: false,
		},
		{name: "string",
			args: args{x: "two"},
			want: true,
		},
		{name: "error",
			args: args{x: errors.New("")},
			want: false,
		},
		{name: "slice",
			args: args{x: []int{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is[string](tt.args.x); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTernary_string(t *testing.T) {
	const evenDay = "even day"
	const oddDay = "odd day"
	timeNowDay := time.Now().Day()
	var res string
	if timeNowDay%2 == 0 {
		res = evenDay
	} else {
		res = oddDay
	}
	type args struct {
		condition bool
		trueT     string
		falseT    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				condition: func() bool {
					return timeNowDay%2 == 0
				}(),
				trueT:  evenDay,
				falseT: oddDay,
			},
			want: res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ternary(tt.args.condition, tt.args.trueT, tt.args.falseT)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

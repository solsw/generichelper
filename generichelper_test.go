package generichelper

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/solsw/builtinhelper"
	"github.com/solsw/mathhelper"
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

func TestMust_panic(t *testing.T) {
	got := func() (err error) {
		defer func() {
			builtinhelper.PanicToError(recover(), &err)
		}()
		Must(23, errors.New("Must error"))
		return nil
	}()
	want := "Must error"
	if !reflect.DeepEqual(got.Error(), want) {
		t.Errorf("Must_panic = %v, want %v", got, want)
	}
}

func TestMust_int(t *testing.T) {
	want := 23
	if got := Must(23, nil); !reflect.DeepEqual(got, want) {
		t.Errorf("Must[int]() = %v, want %v", got, want)
	}
}

func TestTernary_string(t *testing.T) {
	const evenDay = "even day"
	const oddDay = "odd day"
	timeNowDay := time.Now().Day()
	var res string
	if mathhelper.IsEven(int64(timeNowDay)) {
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
					return mathhelper.IsEven(int64(timeNowDay))
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

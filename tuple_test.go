package generichelper

import (
	"reflect"
	"testing"
)

func TestNewTuple2_int_string(t *testing.T) {
	tests := []struct {
		name string
		f    func() (int, string)
		want Tuple2[int, string]
	}{
		{name: "1",
			f:    func() (int, string) { return 1, "one" },
			want: Tuple2[int, string]{1, "one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTuple2(tt.f()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTuple2() = %v, want %v", got, tt.want)
			}
		})
	}
}

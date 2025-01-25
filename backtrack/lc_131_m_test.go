package backtrack

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{name: "test", s: "aab", want: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

package slidingwindow

import "testing"

func Test_countCompleteSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{word: "igigee", k: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countCompleteSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

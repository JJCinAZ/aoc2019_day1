package function

import (
	"testing"
)


func Test_calcFuel(t *testing.T) {
	type args struct {
		mass int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{14}, 2},
		{"test2", args{1969}, 966},
		{"test3", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuel(tt.args.mass); got != tt.want {
				t.Errorf("calcFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
package main

import (
	"testing"
)

func TestToFloat64(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{"12.3333"}, 12.3333},
		{"2", args{"-1.0"}, -1.0},
		{"2.1", args{"1.0"}, 1.0},
		{"3", args{"12.905"}, 12.905},
		{"4", args{"-33.4117"}, -33.4117},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.input); got != tt.want {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64_1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// {"1", args{"12.3333"}, 12.3333},
		// {"2", args{"-1.0"}, -1.0},
		// {"2.1", args{"1.0"}, 1.0},
		// {"3", args{"12.905"}, 12.905},
		{"4", args{"-33.3667"}, -33.3667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64_1(tt.args.input); got != tt.want {
				t.Errorf("ToFloat64_1() = %f, want %f", got, tt.want)
			}
		})
	}
}

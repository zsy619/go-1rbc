package main

import "testing"

func TestToInt32(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 int32
	}{
		{"1", args{"12.3333"}, 123333, 4},
		{"2", args{"-1.0"}, -10, 1},
		{"2.1", args{"1.0"}, 10, 1},
		{"3", args{"12.905"}, 12905, 3},
		{"4", args{"-33.3667"}, -333667, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt32(tt.args.input)
			if got != tt.want {
				t.Errorf("ToInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

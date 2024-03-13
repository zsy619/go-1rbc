package main

import (
	"testing"
)

func TestToCut(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
	}{
		// {"1", args{"Pakhachi;60.5816"}},
		// {"2", args{"Zvëzdnyy;70.9333"}},
		// {"3", args{"Indiga;67.6898"}},
		{"4", args{"Zemlya Bunge;74.8983"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ToCut(tt.args.line)
		})
	}
}

func TestToCut_Split(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{"Pakhachi;6.5816"}},
		{"2", args{"Zvëzdnyy;-7.9333"}},
		{"3", args{"Indiga;67.6898"}},
		{"4", args{"Zemlya Bunge;-74.8983"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ToCut_Split(tt.args.line)
		})
	}
}

package basics_test

import (
	"reflect"
	"testing"

	"github.com/ahnsv/golang-playground/internals/basics"
)

func TestLearnBasics(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
		{name: "basics", want: []string{"hello", "world", "sofam"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basics.LearnBasics(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LearnBasics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLearnFunction(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"basics", args{1}, []string{"hello", "sofam"}},
		{"basics 2", args{0}, []string{"world", "sofam"}},
		{"basics 3", args{2}, []string{"hello", "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basics.LearnFunction(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LearnFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

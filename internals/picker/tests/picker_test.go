package picker_test

import (
	"testing"
	"time"

	"github.com/ahnsv/golang-playground/internals/picker"
)

func TestPickOne(t *testing.T) {
	type args struct {
		s      []string
		srcNum int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Randomly Failing", args{[]string{"hello", "world", "socar"}, time.Now().UnixNano()}, "socar"},
		{"Fixed Source Number", args{[]string{"hello", "world", "socar"}, picker.ParseDateString("2006-01-02", "2021-05-28").UnixNano()}, "world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := picker.PickOne(tt.args.s, int(tt.args.srcNum)); got != tt.want {
				t.Errorf("PickOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

package email

import (
	"testing"
)

func Test_validateTemplate(t *testing.T) {
	tests := []struct {
		val string
		err bool
	}{
		{"example[fn]@gmail.com", false},
		{"e[nws][sws]@a.b", false},
		{"e[nws][sws][fn][by][fs][pby][a]@a.b", false},
		{"e[nws][sws][fn][by][fs][pby][b]@a.b", true},
		{"d[fn{2/3}]", false},
		{"d[nws{e/2}]", false},
		{"d[pby{e4/4}]", false},
		{"d[by{v/3}]", false},
		{"d[fs{ev/2}]", false},
		{"d[fs{ev/}]", true},
		{"d[fs{ve/2}]", true},
		{"d[fs{v4/2}]", true},
		{"d[fs{1231/456}]", false},
	}
	for _, tt := range tests {
		res := validateTemplate(tt.val)
		if (!tt.err && res != nil) || (tt.err && res == nil) {
			t.Errorf(
				"DATA: val: %v EXPECTED ERROR: %v, GOT: %v",
				tt.val,
				tt.err,
				res,
			)
		}
	}
}

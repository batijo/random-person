package utils

import "testing"

func Test_StringContainsInt(t *testing.T) {
	tests := []struct {
		val string
		i   int
		exp bool
	}{
		{"1 2 3 4", 2, true},
		{"1 2 3 s", 2, true},
		{"1 2 2 q", 2, true},
		{"a 2 b c", 2, true},
		{"a 	2 b c", 2, true},
		{" a 2 	 b c", 2, true},
		{"a	2	b c", 2, true},
		{"a 12 b c", 12, true},
		{"a 2bc", 2, false},
		{"a 1 bc", 2, false},
	}
	for _, d := range tests {
		res := StringContainsInt(d.val, d.i)
		if res != d.exp {
			t.Errorf(
				"DATA: val: \"%v\" i: \"%v\" EXPECTED: %v, GOT: %v",
				d.val,
				d.i,
				d.exp,
				res,
			)
		}
	}
}

func Benchmark_StringContainsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringContainsInt("1 2 3 4", 2)
	}
}

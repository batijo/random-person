package utils

import "testing"

func Test_IsCharInString(t *testing.T) {
	tests := []struct {
		val      string
		elements string
		exp      bool
	}{
		{"č", "čšž", true},
		{"Č", "čšž", true},
		{"š", "čšž", true},
		{"ž", "čšž", true},
		{"Ž", "čšž", true},
		{"ą", "čšž", false},
		{"a", "čšž", false},
		{"a", "čašž", true},
		{"!!", "čašž", false},
		{"a", "ASD", true},
		{"as", "asb", false},
	}
	for _, d := range tests {
		res := IsCharInString(d.val, d.elements)
		if res != d.exp {
			t.Errorf(
				"DATA: val: %v elements: %v EXPECTED: %v, GOT: %v",
				d.val,
				d.elements,
				d.exp,
				res,
			)
		}
	}
}

func Test_Trim(t *testing.T) {
	tests := []struct {
		val       string
		elemCount int
		right     bool
		exp       string
	}{
		{"asddqfgq", 1, true, "asddqfg"},
		{"asddqfgq", 1, false, "sddqfgq"},
		{"SAsdaflLP", 2, false, "sdaflLP"},
		{"SAsdaflLP", 2, true, "SAsdafl"},
		{"SAsdaflLP", 100, true, ""},
		{"SAsdaflLP", 100, false, ""},
		{"SAsdaflLP", 0, true, "SAsdaflLP"},
		{"SAsdaflLP", 0, false, "SAsdaflLP"},
		{"SAsdaflLP", -1, true, "SAsdaflLP"},
		{"SAsdaflLP", -1, false, "SAsdaflLP"},
		{"SAsdaflLP", 8, true, "S"},
		{"SAsdaflLP", 8, false, "P"},
		{"", 8, false, ""},
		{"s", 1, false, ""},
		{"", 8, true, ""},
		{"s", 1, true, ""},
		{"s", 0, false, "s"},
		{"s", 0, true, "s"},
		{"čęėčė", 1, false, "ęėčė"},
		{"ČĘĖČĘĖČĖ", 2, false, "ĖČĘĖČĖ"},
		{"ĘĖęėŪūęč", 3, false, "ėŪūęč"},
		{"čęėčė", 1, true, "čęėč"},
		{"ČĘĖČĘĖČĖ", 2, true, "ČĘĖČĘĖ"},
		{"ĘĖęėŪūęč", 3, true, "ĘĖęėŪ"},
	}
	for _, d := range tests {
		res := Trim(d.val, d.elemCount, d.right)
		if res != d.exp {
			t.Errorf(
				"DATA: val: %v elemCount: %v right: %v EXPECTED: %v, GOT: %v",
				d.val, d.elemCount,
				d.right,
				d.exp,
				res,
			)
		}
	}
}

func Test_StrElem(t *testing.T) {
	tests := []struct {
		val  string
		elem int
		exp  string
	}{
		{"sdd", 0, "s"},
		{"SDD", 0, "S"},
		{"sdd", 1, "d"},
		{"SDD", 1, "D"},
		{"SDD", 3, ""},
		{"čėįfš", 4, "š"},
		{"čėįfš", 5, ""},
		{"sdasDdFASasdad", 5, "d"},
	}
	for _, d := range tests {
		res := StrElem(d.val, d.elem)
		if res != d.exp {
			t.Errorf(
				"DATA: val: %v elem: %v EXPECTED: %v, GOT: %v",
				d.val,
				d.elem,
				d.exp,
				res,
			)
		}
	}
}

func Test_FilterChars(t *testing.T) {
	tests := []struct {
		val   string
		chars string
		exp   string
	}{
		{"asdąčęfdąąrągėg", "čąd", "asęfrgėg"},
		{"av564sdčę", "5č", "av64sdę"},
		{"av564sdčęa", "av", "564sdčę"},
		{"av564sdčęa", "0123456789", "avsdčęa"},
		{"avsdčęa", "0123456789", "avsdčęa"},
	}
	for _, d := range tests {
		res := FilterChars(d.val, d.chars)
		if res != d.exp {
			t.Errorf(
				"DATA: val: %v chars: %v EXPECTED: %v, GOT: %v",
				d.val,
				d.chars,
				d.exp,
				res,
			)
		}
	}
}

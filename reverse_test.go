package stringutil

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"hello, 世界", "界世 ,olleh"},
		{"", ""},
	}

	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			// 使用 t.Errorf or t.Fatal
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

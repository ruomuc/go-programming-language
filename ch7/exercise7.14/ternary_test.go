package main

import (
	"fmt"
	"testing"
)

func TestTernary(t *testing.T) {
	var tests = []struct {
		input string
		env   Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"x-9  ? 1 : 2", Env{"x": 9}, "2"},
		{"x-9  ? 1 : 2", Env{"x": 10}, "2"},
	}

	for _, test := range tests {
		expr, err := Parse(test.input)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: got %q, want %q", test.input, err, test.want)
			}
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Println(got)
		if got != test.want {
			t.Errorf("%s: %v => %s, want %s",
				test.input, test.env, got, test.want)
		}
	}
}

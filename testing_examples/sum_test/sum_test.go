package sum_test

import "testing"
import "github.com/iecheniq/go_bootcamp/testing_examples/sum"

func TestSum(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "equals zero", a: 1, b: -1, expected: 0},
		{name: "equals 2", a: 1, b: 1, expected: 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			if res := sum.Sum(c.a, c.b); res != c.expected {
				tt.Errorf("expected %d but got %d", c.expected, res)
			}

		})
	}
}

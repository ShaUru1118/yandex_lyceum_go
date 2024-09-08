package printer

import "testing"

var cases = []struct {
	name   string
	a      int
	b      int
	wanted int
}{
	{
		name:   "first",
		a:      1,
		b:      2,
		wanted: 3,
	},
}

func TestSum(t *testing.T) {

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.a, test.b)
			if got != test.wanted {
				t.Fatalf(`Sum(1, 2) = "%v", want "%v" :/`, got, test.wanted)
			}
		})
	}
}

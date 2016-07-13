package operation

import "testing"

var addCases = []struct {
	x int
	y int
	r int
}{
	{1, 1, 2},
	{20, 21, 41},
	{10, 1000, 1010},
	{8000, 9000, 17000},
	{35, 45, 80},
	{95, 5, 100},
	{3500, 5000, 8500},
	{6, 5, 11},
}

func TestAdd(t *testing.T) {
	for _, a := range addCases {
		res := Add(a.x, a.y)
		if res != a.r {
			t.Errorf("Add(%q, %q) = %q, want %q", a.x, a.y, res, a.r)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add(8000, 9000)
	}
}

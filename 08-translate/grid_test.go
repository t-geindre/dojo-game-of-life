package d08

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func BenchmarkGrid_NextState(b *testing.B) {
	g := NewGrid(2000, 2000)
	g.Randomize(.3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.NextState()
	}
}

func TestGrid_NextState(t *testing.T) {
	tests := []struct {
		name string
		in   []bool
		out  []bool
		it   int
	}{
		{
			name: "Empty world",
			in: []bool{
				false, false, false,
				false, false, false,
				false, false, false,
			},
			out: []bool{
				false, false, false,
				false, false, false,
				false, false, false,
			},
			it: 1,
		},
		{
			name: "Underpopulation",
			in: []bool{
				false, false, false,
				false, true, false,
				false, false, true,
			},
			out: []bool{
				false, false, false,
				false, false, false,
				false, false, false,
			},
			it: 1,
		},
		{
			name: "Stable",
			in: []bool{
				false, false, false, false, false,
				false, true, false, false, false,
				false, false, true, false, false,
				false, false, false, true, false,
				false, false, false, false, false,
			},
			out: []bool{
				false, false, false, false, false,
				false, false, false, false, false,
				false, false, true, false, false,
				false, false, false, false, false,
				false, false, false, false, false,
			},
			it: 1,
		},
		{
			name: "Overpopulation + Reproduction",
			in: []bool{
				false, false, false, false, false,
				false, true, false, true, false,
				false, false, true, false, false,
				false, true, false, true, false,
				false, false, false, false, false,
			},
			out: []bool{
				false, false, false, false, false,
				false, false, true, false, false,
				false, true, false, true, false,
				false, false, true, false, false,
				false, false, false, false, false,
			},
			it: 1,
		},
		{
			name: "Oscillator",
			in: []bool{
				false, false, false, false, false,
				false, false, false, false, false,
				false, true, true, true, false,
				false, false, false, false, false,
				false, false, false, false, false,
			},
			out: []bool{
				false, false, false, false, false,
				false, false, true, false, false,
				false, false, true, false, false,
				false, false, true, false, false,
				false, false, false, false, false,
			},
			it: 3,
		},
		{
			name: "Loop overpopulation",
			in: []bool{
				true, false, true,
				true, false, true,
				true, false, false,
			},
			out: []bool{
				false, false, false,
				false, false, false,
				false, false, false,
			},
			it: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := int(math.Sqrt(float64(len(test.in))))
			g := NewGrid(s, s)
			g.cells = test.in
			for i := test.it; i > 0; i-- {
				g.NextState()
			}
			assert.Equal(t, test.out, g.cells)
		})
	}
}

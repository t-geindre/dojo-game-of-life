package d03

import (
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

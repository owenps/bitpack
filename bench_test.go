package bitpack_test

import (
	"testing"

	"github.com/owenps/bitpack"
)

// Widths span: narrow (3, 5), power-of-two (8), and full (64). Width 5 is
// the interesting case — 64/5 has a remainder, so some indices straddle
// slot boundaries and exercise the two-slot Set/At path.
var benchWidths = []int{3, 5, 8, 64}

const benchN = 1 << 16

func BenchmarkSet(b *testing.B) {
	for _, w := range benchWidths {
		b.Run(name(w), func(b *testing.B) {
			a := bitpack.New(benchN, w)
			v := uint64(1)<<w - 1
			if w == 64 {
				v = ^uint64(0)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				a.Set(i%benchN, v)
			}
		})
	}
}

func BenchmarkAt(b *testing.B) {
	for _, w := range benchWidths {
		b.Run(name(w), func(b *testing.B) {
			a := bitpack.New(benchN, w)
			v := uint64(1)<<w - 1
			if w == 64 {
				v = ^uint64(0)
			}
			for i := range benchN {
				a.Set(i, v)
			}
			b.ResetTimer()
			var sink uint64
			for i := 0; i < b.N; i++ {
				sink = a.At(i % benchN)
			}
			_ = sink
		})
	}
}

// BenchmarkFillVsSetLoop compares Fill to the naive equivalent (calling
// Set in a loop). One op = filling the entire benchN-element array.
func BenchmarkFillVsSetLoop(b *testing.B) {
	for _, w := range benchWidths {
		v := uint64(1)<<w - 1
		if w == 64 {
			v = ^uint64(0)
		}

		b.Run("Fill/"+name(w), func(b *testing.B) {
			a := bitpack.New(benchN, w)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				a.Fill(v)
			}
		})

		b.Run("SetLoop/"+name(w), func(b *testing.B) {
			a := bitpack.New(benchN, w)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for j := 0; j < benchN; j++ {
					a.Set(j, v)
				}
			}
		})
	}
}

func name(width int) string {
	switch width {
	case 3:
		return "width=3"
	case 5:
		return "width=5_crossSlot"
	case 8:
		return "width=8"
	case 64:
		return "width=64"
	}
	return "unknown"
}

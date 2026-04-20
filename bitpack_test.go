package bitpack

import "testing"

func TestSetAndAt(t *testing.T) {
	a := New(5, 3) // 5 elements, 3-bit values (max value = 7)

	values := []uint64{0, 7, 3, 5, 1}
	for i, v := range values {
		a.Set(i, v)
	}

	for i, want := range values {
		got := a.At(i)
		if got != want {
			t.Errorf("Get(%d) = %d, want %d", i, got, want)
		}
	}
}

func TestLenAndBitWidth(t *testing.T) {
	a := New(100, 4)

	if a.Len() != 100 {
		t.Errorf("Len() = %d, want 100", a.Len())
	}
	if a.BitWidth() != 4 {
		t.Errorf("BitWidth() = %d, want 4", a.BitWidth())
	}
}

func TestSize(t *testing.T) {
	// 3-bit values × 100 elements = 300 bits = 4.69 uint64s → 5 uint64s = 40 bytes
	a := New(100, 3)
	if a.Size() != 40 {
		t.Errorf("Size() = %d, want 40", a.Size())
	}

	// 64-bit values × 10 elements = 640 bits = 10 uint64s = 80 bytes
	b := New(10, 64)
	if b.Size() != 80 {
		t.Errorf("Size() = %d, want 80", b.Size())
	}
}

func TestZeroWidth(t *testing.T) {
	a := New(5, 0)

	if a.Len() != 5 {
		t.Errorf("Len() = %d, want 5", a.Len())
	}
	if a.BitWidth() != 0 {
		t.Errorf("BitWidth() = %d, want 0", a.BitWidth())
	}
	if a.Size() != 0 {
		t.Errorf("Size() = %d, want 0", a.Size())
	}

	for i := range 5 {
		a.Set(i, 0)
	}
	for i := range 5 {
		if a.At(i) != 0 {
			t.Errorf("At(%d) = %d, want 0", i, a.At(i))
		}
	}
}

func TestCrossingWordBoundary(t *testing.T) {
	// 5-bit values: 64/5 = 12 values per uint64 with 4 bits left over.
	// The 13th value straddles two uint64s.
	a := New(20, 5)

	for i := range 20 {
		a.Set(i, uint64(i%32)) // max 5-bit value = 31
	}

	for i := range 20 {
		want := uint64(i % 32)
		got := a.At(i)
		if got != want {
			t.Errorf("Get(%d) = %d, want %d", i, got, want)
		}
	}
}

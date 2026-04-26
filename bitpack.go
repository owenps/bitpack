// Package bitpack provides a compact array that stores fixed-width
// bit sequences packed tightly into a flat []uint64.
//
// This is useful when all values fit within a known bit width in the
// range [0, 64], avoiding the space overhead of a standard []uint64
// for narrower values.
package bitpack

import "math"

// Array is a packed array of fixed-width bit sequences.
//
// Multiple goroutines may read from an Array concurrently, but writes
// must be synchronized with all other accesses.
type Array struct {
	data  []uint64
	width uint
	mask  uint64
	size  int
}

// New creates a packed array that stores n values, each using the
// given number of bits. Bit width must be in the range [0, 64].
//
// New panics if the bit width is out of range or if the array size overflows.
func New(n, bitWidth int) *Array {
	if bitWidth < 0 || bitWidth > 64 {
		panic("bitpack: bit width must be in range [0, 64]")
	}

	if n < 0 {
		panic("bitpack: number of values must be non-negative")
	}

	if bitWidth > 0 && (math.MaxInt-63)/bitWidth < n {
		panic("bitpack: array size overflows int")
	}

	// Number of uint64 slots needed, rounded up.
	slots := (n*bitWidth + 63) / 64

	return &Array{
		data:  make([]uint64, slots),
		width: uint(bitWidth),
		mask:  ^uint64(0) >> (64 - bitWidth),
		size:  n,
	}
}

// Set stores a value at the given index.
//
// Set panics if the index is out of range or if the value exceeds the bit width.
func (a *Array) Set(i int, value uint64) {
	if i < 0 || i >= a.size {
		panic("bitpack: index out of range")
	}

	if value>>a.width != 0 {
		panic("bitpack: value exceeds bit width")
	}

	if a.width == 0 {
		return
	}

	bitPos := i * int(a.width)
	slot := bitPos / 64
	offset := uint(bitPos % 64)

	if offset+a.width <= 64 {
		a.data[slot] = (a.data[slot] &^ (a.mask << offset)) | (value << offset)
		return
	}

	// The value spans multiple slots; we need to write to two slots.
	loBits := value << offset
	hiBits := value >> (64 - offset)
	a.data[slot] = (a.data[slot] &^ (a.mask << offset)) | loBits
	a.data[slot+1] = (a.data[slot+1] &^ (a.mask >> (64 - offset))) | hiBits
}

// At returns the value at the given index.
//
// At panics if the index is out of range.
func (a *Array) At(i int) uint64 {
	if i < 0 || i >= a.size {
		panic("bitpack: index out of range")
	}

	if a.width == 0 {
		return 0
	}

	bitPos := i * int(a.width)
	slot := bitPos / 64
	offset := uint(bitPos % 64)

	if offset+a.width <= 64 {
		return (a.data[slot] >> offset) & a.mask
	}

	// The value spans multiple slots; we need to read from two slots.
	loBits := a.data[slot] >> offset
	hiBits := a.data[slot+1] << (64 - offset)
	return (loBits | hiBits) & a.mask
}

// Fill sets all elements in the array to the given value.
//
// Fill panics if the value exceeds the bit width.
func (a *Array) Fill(value uint64) {
	if value>>a.width != 0 {
		panic("bitpack: value exceeds the bit width.")
	}

	if a.width == 0 {
		return
	}

	// Precompute a repeated pattern of the value for filling the array.
	stamp := value
	shift := a.width
	for shift < 64 {
		stamp = (stamp << shift) | stamp
		shift = shift * 2
	}

	rotate := 64 % a.width
	period := min(a.width/gcd(a.width, rotate), uint(len(a.data)))
	for i := range period {
		a.data[i] = stamp
		stamp = (stamp << rotate) | (stamp >> (64 - rotate))
	}

	filled := period
	for filled < uint(len(a.data)) {
		copy(a.data[filled:], a.data[:filled])
		filled *= 2
	}
}

// Len returns the number of elements in the array.
func (a *Array) Len() int {
	return a.size
}

// BitWidth returns the number of bits used per element.
func (a *Array) BitWidth() int {
	return int(a.width)
}

// Size returns the number of bytes used by the underlying storage.
func (a *Array) Size() int {
	return len(a.data) * 8
}

// gcd returns the greatest common divisor of two integers.
func gcd(a, b uint) uint {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

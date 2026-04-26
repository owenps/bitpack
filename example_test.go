package bitpack_test

import (
	"fmt"

	"github.com/owenps/bitpack"
)

func Example() {
	a := bitpack.New(8, 3)
	for i := range 8 {
		a.Set(i, uint64(i))
	}

	for i := range a.Len() {
		fmt.Println(a.At(i))
	}

	fmt.Printf("stored %d values in %d bytes\n", a.Len(), a.Size())

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// stored 8 values in 8 bytes
}

func ExampleFromSlice() {
	a := bitpack.FromSlice([]uint64{0, 1, 2, 3, 4, 5, 6, 7})
	for i := range a.Len() {
		fmt.Println(a.At(i))
	}

	fmt.Printf("stored %d values in %d bytes\n", a.Len(), a.Size())

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// stored 8 values in 8 bytes
}

func ExampleArray_ToSlice() {
	a := bitpack.FromSlice([]uint64{0, 1, 2, 3, 4, 5, 6, 7})
	slice := a.ToSlice()
	for i, v := range slice {
		fmt.Printf("slice[%d] = %d\n", i, v)
	}

	// Output:
	// slice[0] = 0
	// slice[1] = 1
	// slice[2] = 2
	// slice[3] = 3
	// slice[4] = 4
	// slice[5] = 5
	// slice[6] = 6
	// slice[7] = 7
}

func ExampleArray_Fill() {
	a := bitpack.FromSlice([]uint64{0, 1, 2, 3, 4, 5, 6, 7})
	a.Fill(1)
	for i := range a.Len() {
		fmt.Println(a.At(i))
	}

	// Output:
	// 1
	// 1
	// 1
	// 1
	// 1
	// 1
	// 1
	// 1
}

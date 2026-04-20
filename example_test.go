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

# bitpack ⚡︎

[![CI](https://github.com/owenps/bitpack/actions/workflows/test.yml/badge.svg)](https://github.com/owenps/bitpack/actions/workflows/test.yml)
[![Version](https://img.shields.io/github/v/release/owenps/bitpack)](https://github.com/owenps/bitpack/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/owenps/bitpack.svg)](https://pkg.go.dev/github.com/owenps/bitpack)
[![Go Report Card](https://goreportcard.com/badge/github.com/owenps/bitpack)](https://goreportcard.com/report/github.com/owenps/bitpack)
[![License](https://img.shields.io/github/license/owenps/bitpack)](/LICENSE)

A miniature Go library for compact arrays.

> [!NOTE]
> Read more of my developer notes from [my blog](https://owenps.github.io/blog).

## Usage

```go
package main

import (
	"fmt"
	"github.com/owenps/bitpack"
)

func main() {
	a := bitpack.New(1000, 3) // 1000 elements, 3 bits per element.

	for i := range a.Len() {
		a.Set(i, uint64(i%8)) // Set value from [0, 7]
	}

	fmt.Printf("stored %d values in %d bytes\n", a.Len(), a.Size())
	fmt.Printf("[]uint64 would need %d bytes\n", a.Len()*8)
}
```

```text
stored 1000 values in 376 bytes
[]uint64 would need 8000 bytes
```

For more examples see [example_test.go](/example_test.go)

## Benchmarks

All operations run in constant time with zero allocations.
Single-digit nanoseconds per `Set`/`At` call on modern hardware.
Run `go test -bench=. -benchmem` to measure on your own machine.

## Installation

```bash
go get github.com/owenps/bitpack
```

## Road Map

- [ ] `FromSlice()` method - automatically calculate the width from a slice.
- [ ] Generics support - support other envelopes other than `uint64` (`uint8`, `uint32`, `uint`)

## License

[MIT](/LICENSE)

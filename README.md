# BitPack ⚡︎

![github.com/owenps/bitpack/actions/workflows/test.yml/badge.svg](https://github.com/owenps/bitpack/actions/workflows/test.yml/badge.svg)
![pkg.go.dev/badge/github.com/owenps/bitpack.svg](https://pkg.go.dev/badge/github.com/owenps/bitpack.svg)
![goreportcard.com/badge/github.com/owenps/bitpack](https://goreportcard.com/badge/github.com/owenps/bitpack)
![img.shields.io/github/license/owenps/bitpack](https://img.shields.io/github/license/owenps/bitpack)

A miniature GO library for compact arrays.

> [!NOTE]
> Read more of my developer notes from [my blog](https://owenps.github.io/blog).

## Usage

```go
import "github.com/owenps/bitpack"

x := uint64{0, 6, 1, 6} // largest value is 6 (bit width of 3). 

y := bitpack.New(4, 3)  // 4 elements, 3 bits per element.

y.Set(0, 0)
y.Set(1, 6)
y.Set(2, 1)
y.Set(3, 6)

for i := range y.Len() {
	fmt.Println(y.At(i))
}
```

For more examples see [example_test.go](/example_test)

## Benchmarks



Run tests with 

```bash
go test
```

## Installation

```bash
go get github.com/owenps/bitpack 
```

## Road Map

- [ ] `FromSlice()` method - automatically calculate the width from a slice.
- [ ] Generics support - support other envelopes other than (`uint8`, `uint32`, `uint`)

## License

[MIT](/LICENSE)

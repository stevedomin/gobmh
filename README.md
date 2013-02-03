gobmh
=====

Go implementation of Boyer-Moore-Horspool algorithm.

## Install

To use it in your project :

```go
import (
	"github.com/stevedomin/gobmh"
)
```

## Benchmarks

Run the benchmarks with :

```bash
$ go test -test.bench="BenchmarkIndex"
```

On my Intel Core i7 2.6Ghz, the BMH algorithm is 3-4 times faster than bytes.Index, depending on the size of the "haystack".  
I've tested against "History of Western Philosophy", from Bertrand Russel.

```
BenchmarkIndex                  500			6446828 ns/op  
BenchmarkIndexHorspool         1000			2126440 ns/op
```

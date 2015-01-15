Introduction
============
Some little experiments for bytes package.

Benchmark
=========

```
Benchmark_buffer_cap     20000000                66.1 ns/op      15498.40 MB/s         31 B/op          0 allocs/op
--- BENCH: Benchmark_buffer_cap
main_test.go:19: 64 11
main_test.go:19: 1189 1100
main_test.go:19: 153589 110000
main_test.go:19: 19660789 11000000
main_test.go:19: 314572789 220000000
ok      golabs/bytes    1.465s
```
From the benchmark results above, buffer cannot set its cap to 0,
which may lead to memory leak.

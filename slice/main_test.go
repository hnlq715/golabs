package main

import "testing"

// bytes.Buffer has no method to clean all the memory allocated.
func Benchmark_buffer_cap(b *testing.B) {
	b.SetBytes(1024)
	b.StopTimer()
	b.StartTimer()

	s := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		s = append(s, 's')

	}
	s = nil

	b.Log(cap(s), len(s))
}

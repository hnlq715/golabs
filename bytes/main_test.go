package main

import (
	"bytes"
	"runtime"
	"testing"
)

// bytes.Buffer has no method to clean all the memory allocated.
func Benchmark_buffer_cap(b *testing.B) {
	b.SetBytes(1024)
	b.StopTimer()
	b.StartTimer()

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.WriteString("Hello world")
	}

	b.Log(cap(buf.Bytes()), len(buf.Bytes()))
}

// bytes.Buffer has no method to clean all the memory allocated.
func Benchmark_buffer_gc(b *testing.B) {
	b.SetBytes(1024)
	b.StopTimer()
	b.StartTimer()

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.WriteString("Hello world")
	}

	b.Log(cap(buf.Bytes()), len(buf.Bytes()))
	runtime.GC()

}

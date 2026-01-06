package main

import "testing"

func BenchmarkGoToRust(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rustAdd(1, 2)
	}
}

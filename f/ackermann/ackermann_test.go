package main

import "testing"

func benchmarkAck(m int, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		ackermann(m, n)
	}
}

func BenchmarkAck3_7(b *testing.B) {
	benchmarkAck(3, 7, b)
}

func BenchmarkAck3_8(b *testing.B) {
	benchmarkAck(3, 8, b)
}

func BenchmarkAck3_9(b *testing.B) {
	benchmarkAck(3, 9, b)
}

func BenchmarkAck3_10(b *testing.B) {
	benchmarkAck(3, 10, b)
}

func BenchmarkAck3_11(b *testing.B) {
	benchmarkAck(3, 11, b)
}

func BenchmarkAck3_12(b *testing.B) {
	benchmarkAck(3, 12, b)
}

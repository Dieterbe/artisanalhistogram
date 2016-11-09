package hist1

import (
	"math/rand"
	"testing"
	"time"
)

// all values under 1ms so they go into first bucket
func Benchmark_AddDurationBest(b *testing.B) {
	data := make([]time.Duration, b.N)
	hist := New()
	for i := 0; i < b.N; i++ {
		data[i] = time.Duration(rand.Intn(1000)) * time.Microsecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hist.AddDuration(data[i])
	}
}

// all values over 15s so they go into last bucket
func Benchmark_AddDurationWorst(b *testing.B) {
	data := make([]time.Duration, b.N)
	hist := New()
	for i := 0; i < b.N; i++ {
		data[i] = time.Duration(16+rand.Intn(10)) * time.Second
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hist.AddDuration(data[i])
	}
}

// all between 0ms and 20s to they go anywhere. but later buckets get higher proportion cause they cover more ground
func Benchmark_AddDurationEvenDistribution(b *testing.B) {
	data := make([]time.Duration, b.N)
	hist := New()
	for i := 0; i < b.N; i++ {
		data[i] = time.Duration(rand.Intn(20000000)) * time.Microsecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hist.AddDuration(data[i])
	}
}

// all between 0ms and 1s. more realistic. control over distribution would be better though
func Benchmark_AddDurationUpto1s(b *testing.B) {
	data := make([]time.Duration, b.N)
	hist := New()
	for i := 0; i < b.N; i++ {
		data[i] = time.Duration(rand.Intn(1000)) * time.Millisecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hist.AddDuration(data[i])
	}
}

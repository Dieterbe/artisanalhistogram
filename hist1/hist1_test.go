package hist1

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Report(t *testing.T) {
	hist := New()
	hist.AddDuration(time.Duration(10) * time.Microsecond)
	hist.AddDuration(time.Duration(4) * time.Millisecond)
	hist.AddDuration(time.Duration(5) * time.Millisecond)
	hist.AddDuration(time.Duration(10) * time.Millisecond)
	hist.AddDuration(time.Duration(1000) * time.Millisecond)
	hist.AddDuration(time.Duration(1000) * time.Millisecond)
	hist.AddDuration(time.Duration(1000) * time.Millisecond)
	hist.AddDuration(time.Duration(1001) * time.Millisecond)
	hist.AddDuration(time.Duration(1200) * time.Millisecond)
	hist.AddDuration(time.Duration(21) * time.Second)

	snap := hist.Snapshot()
	exp := []uint32{
		1, //1000 micros,
		0, //2000,
		0, //3000,
		2, //5000,
		0, //7500,
		1, //10000,
		0, //15000,
		0, //20000,
		0, //30000,
		0, //40000,
		0, //50000,
		0, //65000,
		0, //80000,
		0, //100000,
		0, //150000,
		0, //200000,
		0, //300000,
		0, //400000,
		0, //500000,
		0, //650000,
		0, //800000,
		3, //1000000,
		2, //1500000,
		0, //2000000,
		0, //3000000,
		0, //4000000,
		0, //5000000,
		0, //6500000,
		0, //8000000,
		0, //10000000,
		0, //15000000,
		1, //29999999, // used to represent inf
	}
	for i, cnt := range snap {
		if cnt != exp[i] {
			t.Fatalf("expected snap[%d] = %d, got %d", i, exp[i], cnt)
		}
	}

	r := hist.Report(snap)

	actualTotal := uint32(10 + 4000 + 5000 + 10000 + 3000*1000 + 2002*1000 + 1200000 + 21000000)
	actualMean := actualTotal / 10
	expTotal := uint32(1000 + 2*5000 + 10000 + 3*1000000 + 2*1500000 + 29999999)

	if r.min != 1000 {
		t.Fatalf("expected min %d, got %d", 1000, r.min)
	}

	expMean := expTotal / 10
	if r.mean != expMean {
		t.Fatalf("expected mean %d, got %d (actual mean %d)", expMean, r.mean, actualMean)
	}
	t.Logf("actual mean %d, our mean %d (big outlier!)", actualMean, r.mean)

	if r.median != 1000000 {
		t.Fatalf("expected med %d, got %d", 1000000, r.median)
	}

	if r.p75 != 1500000 {
		t.Fatalf("expected p75 %d, got %d", 1500000, r.p75)
	}

	if r.p90 != 1500000 {
		t.Fatalf("expected min %d, got %d", 1500000, r.p90)
	}

	expMax := uint32(29999999)
	if r.max != expMax {
		t.Fatalf("expected max %d, got %d", expMax, r.max)
	}

	if r.count != 10 {
		t.Fatalf("expected count %d, got %d (actual count %d)", 10, r.count)
	}

}

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

var _report Report

func Benchmark_Report1kvals(b *testing.B) {
	data := make([]time.Duration, 1000)
	hist := New()
	for i := 0; i < 1000; i++ {
		data[i] = time.Duration(rand.Intn(1000)) * time.Millisecond
	}
	var r Report

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		//snapshots resets the state, so we have to repopulate it
		b.StopTimer()
		for i := 0; i < 1000; i++ {
			hist.AddDuration(data[i])
		}
		b.StartTimer()

		snap := hist.Snapshot()
		r = hist.Report(snap)
	}
	_report = r

}

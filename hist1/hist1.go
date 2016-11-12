package hist1

import (
	"math"
	"sync/atomic"
	"time"
)

const maxVal = uint32(29999999) // used to report max number as 29s even if it's higher

type Hist1 struct {
	limits [32]uint32 // in micros
	counts [32]uint32
}

func New() Hist1 {
	return Hist1{
		limits: [32]uint32{
			1000,           // 0
			2000,           // 1
			3000,           // 2
			5000,           // 3
			7500,           // 4
			10000,          // 5
			15000,          // 6
			20000,          // 7
			30000,          // 8
			40000,          // 9
			50000,          // 10
			65000,          // 11
			80000,          // 12
			100000,         // 13
			150000,         // 14
			200000,         // 15
			300000,         // 16
			400000,         // 17
			500000,         // 18
			650000,         // 19
			800000,         // 20
			1000000,        // 21
			1500000,        // 22
			2000000,        // 23
			3000000,        // 24
			4000000,        // 25
			5000000,        // 26
			6500000,        // 27
			8000000,        // 28
			10000000,       // 29
			15000000,       // 30
			math.MaxUint32, // 31 // to ease binary search, but will be reported as 29s
		},
	}
}

// searchBucket implements a binary search, to find the bucket i to insert val in, like so:
// limits[i-1] < val <= limits[i]
// if we can convince the go compiler to inline this we can get a 14~22% speedup (verified by manually patching it in)
// but we can't :( see https://github.com/golang/go/issues/17566
// so for now, we just replicate this code in addDuration below. make sure to keep the code in sync!
func searchBucket(limits [32]uint32, micros uint32) int {
	min, i, max := 0, 16, 32
	for {
		if micros <= limits[i] {
			if i == 0 || micros > limits[i-1] {
				return i
			}
			max = i
		} else {
			min = i
		}
		i = min + ((max - min) / 2)
	}
}

// adds to the right bucket with a copy of the searchBucket function below, to enforce inlining.
func (h *Hist1) AddDuration(value time.Duration) {
	// note: overflows at 4294s, but if you have values this high,
	// you are definitely not using this histogram for the target use case.
	micros := uint32(value.Nanoseconds() / 1000)
	min, i, max := 0, 16, 32
	for {
		if micros <= h.limits[i] {
			if i == 0 || micros > h.limits[i-1] {
				atomic.AddUint32(&h.counts[i], 1)
				return
			}
			max = i
		} else {
			min = i
		}
		i = min + ((max - min) / 2)
	}
}

// Snapshot returns a snapshot of the data and resets internal state
func (h *Hist1) Snapshot() []uint32 {
	snap := make([]uint32, 32)
	for i := 0; i < 32; i++ {
		snap[i] = atomic.SwapUint32(&h.counts[i], 0)
	}
	return snap
}

// if count is 0 then the statistical summaries are invalid
type Report struct {
	min    uint32 // in micros
	mean   uint32 // in micros
	median uint32 // in micros
	p75    uint32 // in micros
	p90    uint32 // in micros
	max    uint32 // in micros
	count  uint64
}

func (h *Hist1) Report(data []uint32) Report {
	totalValue := uint64(0)
	r := Report{}
	for i, count := range data {
		if count > 0 {
			limit := h.limits[i]
			// special case, report as 29s
			if i == 31 {
				limit = maxVal
			}
			if r.min == 0 { // this means we haven't found min yet.
				r.min = limit
			}
			r.max = limit
			r.count += uint64(count)
			totalValue += uint64(count) * uint64(limit)
		}
	}
	if r.count == 0 {
		return r
	}
	r.median = h.limits[Quantile(data, 0.50, r.count)]
	if r.median == math.MaxUint32 {
		r.median = maxVal
	}
	r.p75 = h.limits[Quantile(data, 0.75, r.count)]
	if r.p75 == math.MaxUint32 {
		r.p75 = maxVal
	}
	r.p90 = h.limits[Quantile(data, 0.90, r.count)]
	if r.p90 == math.MaxUint32 {
		r.p90 = maxVal
	}
	r.mean = uint32(totalValue / r.count)
	return r
}

// quantile q means what's the value v so that all q of the values have value <= v
func Quantile(data []uint32, q float64, total uint64) int {
	count := q * float64(total)
	for i := 0; i < 32; i++ {
		count -= float64(data[i])

		if count <= 0 {
			return i // we return the upper limit, real quantile would be less, but we can't make the result any better.
		}
	}

	return -1
}

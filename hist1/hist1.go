package hist1

import (
	"sync/atomic"
	"time"
)

// in micros
var buckets = []int{
	1000,
	2000,
	3000,
	5000,
	7500,
	10000,
	15000,
	20000,
	30000,
	40000,
	50000,
	65000,
	80000,
	100000,
	150000,
	200000,
	300000,
	400000,
	500000,
	650000,
	800000,
	1000000,
	1500000,
	2000000,
	3000000,
	4000000,
	5000000,
	6500000,
	8000000,
	10000000,
	15000000,
	29999999, // used to represent inf
}

type Hist1 struct {
	counts [32]uint32
}

func New() Hist1 {
	return Hist1{}
}

func (h *Hist1) AddDuration(value time.Duration) {
	if value <= time.Duration(1)*time.Millisecond {
		atomic.AddUint32(&h.counts[0], 1)
	} else if value <= time.Duration(2)*time.Millisecond {
		atomic.AddUint32(&h.counts[1], 1)
	} else if value <= time.Duration(3)*time.Millisecond {
		atomic.AddUint32(&h.counts[2], 1)
	} else if value <= time.Duration(5)*time.Millisecond {
		atomic.AddUint32(&h.counts[3], 1)
	} else if value <= time.Duration(7500)*time.Microsecond {
		atomic.AddUint32(&h.counts[4], 1)
	} else if value <= time.Duration(10)*time.Millisecond {
		atomic.AddUint32(&h.counts[5], 1)
	} else if value <= time.Duration(15)*time.Millisecond {
		atomic.AddUint32(&h.counts[6], 1)
	} else if value <= time.Duration(20)*time.Millisecond {
		atomic.AddUint32(&h.counts[7], 1)
	} else if value <= time.Duration(30)*time.Millisecond {
		atomic.AddUint32(&h.counts[8], 1)
	} else if value <= time.Duration(40)*time.Millisecond {
		atomic.AddUint32(&h.counts[9], 1)
	} else if value <= time.Duration(50)*time.Millisecond {
		atomic.AddUint32(&h.counts[10], 1)
	} else if value <= time.Duration(65)*time.Millisecond {
		atomic.AddUint32(&h.counts[11], 1)
	} else if value <= time.Duration(80)*time.Millisecond {
		atomic.AddUint32(&h.counts[12], 1)
	} else if value <= time.Duration(100)*time.Millisecond {
		atomic.AddUint32(&h.counts[13], 1)
	} else if value <= time.Duration(150)*time.Millisecond {
		atomic.AddUint32(&h.counts[14], 1)
	} else if value <= time.Duration(200)*time.Millisecond {
		atomic.AddUint32(&h.counts[15], 1)
	} else if value <= time.Duration(300)*time.Millisecond {
		atomic.AddUint32(&h.counts[16], 1)
	} else if value <= time.Duration(400)*time.Millisecond {
		atomic.AddUint32(&h.counts[17], 1)
	} else if value <= time.Duration(500)*time.Millisecond {
		atomic.AddUint32(&h.counts[18], 1)
	} else if value <= time.Duration(650)*time.Millisecond {
		atomic.AddUint32(&h.counts[19], 1)
	} else if value <= time.Duration(800)*time.Millisecond {
		atomic.AddUint32(&h.counts[20], 1)
	} else if value <= time.Duration(1000)*time.Millisecond {
		atomic.AddUint32(&h.counts[21], 1)
	} else if value <= time.Duration(1500)*time.Millisecond {
		atomic.AddUint32(&h.counts[22], 1)
	} else if value <= time.Duration(2000)*time.Millisecond {
		atomic.AddUint32(&h.counts[23], 1)
	} else if value <= time.Duration(3000)*time.Millisecond {
		atomic.AddUint32(&h.counts[24], 1)
	} else if value <= time.Duration(4000)*time.Millisecond {
		atomic.AddUint32(&h.counts[25], 1)
	} else if value <= time.Duration(5000)*time.Millisecond {
		atomic.AddUint32(&h.counts[26], 1)
	} else if value <= time.Duration(6500)*time.Millisecond {
		atomic.AddUint32(&h.counts[27], 1)
	} else if value <= time.Duration(8)*time.Second {
		atomic.AddUint32(&h.counts[28], 1)
	} else if value <= time.Duration(10)*time.Second {
		atomic.AddUint32(&h.counts[29], 1)
	} else if value <= time.Duration(15)*time.Second {
		atomic.AddUint32(&h.counts[30], 1)
	} else {
		atomic.AddUint32(&h.counts[31], 1)
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

func Report(data []uint32) []int {
	totalCount := uint64(0)
	totalValue := uint64(0)
	var min, max int
	for i, count := range data {
		if count > 0 {
			if min == 0 { // this means we haven't found min yet.
				min = buckets[i]
			}
			max = buckets[i]
			totalCount += uint64(count)
			totalValue += uint64(count) * uint64(buckets[i])
		}
	}
	if totalCount == 0 {
		return nil
	}
	median := buckets[Quantile(data, 0.50, totalCount)]
	p75 := buckets[Quantile(data, 0.75, totalCount)]
	p90 := buckets[Quantile(data, 0.90, totalCount)]
	mean := int(totalValue / totalCount)
	res := []int{min, mean, median, p75, p90, max, int(totalCount)}
	return res
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

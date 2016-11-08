package main

import (
	"fmt"
	"math"
)

type Hist2 struct {
	counts [16]uint32
}

func NewHist2() {
	return Hist2{}
}

/* buckets:
2,
5,
10,
20,
40,
70,
100,
200,
400,
700,
1000,
2000,
4000,
7000,
10000,
inf,
*/

func (h *Hist2) AddDuration(value time.Duration) {
	if value <= time.Duration(2)*time.Millisecond {
		h.counts[0] += 1
	} else if value <= time.Duration(5)*time.Millisecond {
		h.counts[1] += 1
	} else if value <= time.Duration(10)*time.Millisecond {
		h.counts[2] += 1
	} else if value <= time.Duration(20)*time.Millisecond {
		h.counts[3] += 1
	} else if value <= time.Duration(40)*time.Millisecond {
		h.counts[4] += 1
	} else if value <= time.Duration(70)*time.Millisecond {
		h.counts[5] += 1
	} else if value <= time.Duration(100)*time.Millisecond {
		h.counts[6] += 1
	} else if value <= time.Duration(200)*time.Millisecond {
		h.counts[7] += 1
	} else if value <= time.Duration(400)*time.Millisecond {
		h.counts[8] += 1
	} else if value <= time.Duration(700)*time.Millisecond {
		h.counts[9] += 1
	} else if value <= time.Duration(1000)*time.Millisecond {
		h.counts[10] += 1
	} else if value <= time.Duration(2000)*time.Millisecond {
		h.counts[11] += 1
	} else if value <= time.Duration(4000)*time.Millisecond {
		h.counts[12] += 1
	} else if value <= time.Duration(7000)*time.Millisecond {
		h.counts[13] += 1
	} else if value <= time.Duration(10000)*time.Millisecond {
		h.counts[14] += 1
	} else {
		h.counts[15] += 1
	}
}
func (h *Hist2) Report() []uint32 {
	foundMin := false
	foundMax := false
	totalCount := 0
	for i, count := range h.counts {
		if count > 0 {
			if !foundMin {
				min = i
				foundMin = true
			}
			max = i
			foundMax = true
		}
		totalCount += count
	}
	if total == 0 {
		return nil
	}
	// TODO: return min, max, median, mean, percentiles, etc
}

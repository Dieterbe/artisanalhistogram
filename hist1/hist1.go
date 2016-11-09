package hist1

import (
	"sync/atomic"
	"time"
)

type Hist1 struct {
	counts [32]uint32
}

func New() Hist1 {
	return Hist1{}
}

func (h *Hist1) AddDuration(value time.Duration) {
	if value <= time.Duration(1)*time.Millisecond {
		//h.counts[0] += 1
		atomic.AddUint32(&h.counts[0], 1)
	} else if value <= time.Duration(2)*time.Millisecond {
		//h.counts[1] += 1
		atomic.AddUint32(&h.counts[1], 1)
	} else if value <= time.Duration(3)*time.Millisecond {
		//h.counts[2] += 1
		atomic.AddUint32(&h.counts[2], 1)
	} else if value <= time.Duration(5)*time.Millisecond {
		//h.counts[3] += 1
		atomic.AddUint32(&h.counts[3], 1)
	} else if value <= time.Duration(7500)*time.Microsecond {
		//h.counts[4] += 1
		atomic.AddUint32(&h.counts[4], 1)
	} else if value <= time.Duration(10)*time.Millisecond {
		//h.counts[5] += 1
		atomic.AddUint32(&h.counts[5], 1)
	} else if value <= time.Duration(15)*time.Millisecond {
		//h.counts[6] += 1
		atomic.AddUint32(&h.counts[6], 1)
	} else if value <= time.Duration(20)*time.Millisecond {
		//h.counts[7] += 1
		atomic.AddUint32(&h.counts[7], 1)
	} else if value <= time.Duration(30)*time.Millisecond {
		//h.counts[8] += 1
		atomic.AddUint32(&h.counts[8], 1)
	} else if value <= time.Duration(40)*time.Millisecond {
		//h.counts[9] += 1
		atomic.AddUint32(&h.counts[9], 1)
	} else if value <= time.Duration(50)*time.Millisecond {
		//h.counts[10] += 1
		atomic.AddUint32(&h.counts[10], 1)
	} else if value <= time.Duration(65)*time.Millisecond {
		//h.counts[11] += 1
		atomic.AddUint32(&h.counts[11], 1)
	} else if value <= time.Duration(80)*time.Millisecond {
		//h.counts[12] += 1
		atomic.AddUint32(&h.counts[12], 1)
	} else if value <= time.Duration(100)*time.Millisecond {
		//h.counts[13] += 1
		atomic.AddUint32(&h.counts[13], 1)
	} else if value <= time.Duration(150)*time.Millisecond {
		//h.counts[14] += 1
		atomic.AddUint32(&h.counts[14], 1)
	} else if value <= time.Duration(200)*time.Millisecond {
		//h.counts[15] += 1
		atomic.AddUint32(&h.counts[15], 1)
	} else if value <= time.Duration(300)*time.Millisecond {
		//h.counts[16] += 1
		atomic.AddUint32(&h.counts[16], 1)
	} else if value <= time.Duration(400)*time.Millisecond {
		//h.counts[17] += 1
		atomic.AddUint32(&h.counts[17], 1)
	} else if value <= time.Duration(500)*time.Millisecond {
		//h.counts[18] += 1
		atomic.AddUint32(&h.counts[18], 1)
	} else if value <= time.Duration(650)*time.Millisecond {
		//h.counts[19] += 1
		atomic.AddUint32(&h.counts[19], 1)
	} else if value <= time.Duration(800)*time.Millisecond {
		//h.counts[20] += 1
		atomic.AddUint32(&h.counts[20], 1)
	} else if value <= time.Duration(1000)*time.Millisecond {
		//h.counts[21] += 1
		atomic.AddUint32(&h.counts[21], 1)
	} else if value <= time.Duration(1500)*time.Millisecond {
		//h.counts[22] += 1
		atomic.AddUint32(&h.counts[22], 1)
	} else if value <= time.Duration(2000)*time.Millisecond {
		//h.counts[23] += 1
		atomic.AddUint32(&h.counts[23], 1)
	} else if value <= time.Duration(3000)*time.Millisecond {
		//h.counts[24] += 1
		atomic.AddUint32(&h.counts[24], 1)
	} else if value <= time.Duration(4000)*time.Millisecond {
		//h.counts[25] += 1
		atomic.AddUint32(&h.counts[25], 1)
	} else if value <= time.Duration(5000)*time.Millisecond {
		//h.counts[26] += 1
		atomic.AddUint32(&h.counts[26], 1)
	} else if value <= time.Duration(6500)*time.Millisecond {
		//h.counts[27] += 1
		atomic.AddUint32(&h.counts[27], 1)
	} else if value <= time.Duration(8)*time.Second {
		//h.counts[28] += 1
		atomic.AddUint32(&h.counts[28], 1)
	} else if value <= time.Duration(10)*time.Second {
		//h.counts[29] += 1
		atomic.AddUint32(&h.counts[29], 1)
	} else if value <= time.Duration(15)*time.Second {
		//h.counts[30] += 1
		atomic.AddUint32(&h.counts[30], 1)
	} else {
		//h.counts[31] += 1
		atomic.AddUint32(&h.counts[31], 1)
	}
}

/*
func (h *Hist1) Report() []uint32 {
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
*/

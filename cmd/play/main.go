package main

import (
	"fmt"
	"math"
)

// to achieve the goals from the README, here are the implementation tricks
// * math.MaxUint32, i.e. 4294967295 should be a reasonable count limit. to save space.
//   (it's easy to warn operators if they get too close to the limit and make them upgrade to bigger structures
//   or shorter their reporting interval, and we can invalidate numbers if we detect an almost overflow)
// * proposal2 limits the size of the histogram to 64B aka 1 cacheline. i have no idea if it makes sense.

// in ms
// fairly well spread out and good balance between error % and bucket understandability
// but 38 buckets at 4B => 152B
var proposal1 = []uint32{
	1,
	2,
	3,
	4,
	5,
	6,
	8,
	10,
	14,
	18,
	24,
	30,
	40,
	50,
	60,
	80,
	100,
	140,
	180,
	240,
	300,
	400,
	500,
	600,
	800,
	1000,
	1400,
	1800,
	2400,
	3000,
	4000,
	5000,
	6000,
	8000,
	10000,
	15000,
	20000,
	0,
}

// in ms
// fairly well spread out and good balance between error % and bucket understandability
// aggresive reduction of # buckets -> 16 buckets at 4B, fits in a 64B cacheline
// even this should be practical enough for operator
var proposal2 = []uint32{
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
	0,
}

func main() {
	info(proposal1)
	info(proposal2)
}

func info(buckets []uint32) {
	for i, bucket := range buckets {
		prev := uint32(0)
		if i != 0 {
			prev = buckets[i-1]
		}
		if bucket == 0 {
			bucket = math.MaxUint32
		}
		width := bucket - prev
		err := width * 100 / bucket
		fmt.Printf("bucket %10d. diff %d, max error: %d %%\n", bucket, width, err)
	}
	fmt.Println("#buckets:", len(buckets))
}

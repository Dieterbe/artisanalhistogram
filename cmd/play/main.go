package main

import (
	"fmt"
	"math"
)

func main() {
	info(hist1)
	info(proposal1)
	info(proposal3)
	info(proposal5)
}

func info(buckets []uint32) {
	totalErr := uint32(0)
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
		totalErr += err
		fmt.Printf("bucket %10d. diff %d, max error: %d %%\n", bucket, width, err)
	}
	fmt.Println("#buckets:", len(buckets), "total err", totalErr)
}

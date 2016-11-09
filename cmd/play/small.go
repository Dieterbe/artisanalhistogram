package main

// in ms
// fairly well spread out and good balance between error % and bucket understandability
// aggresive reduction of # buckets -> 16 buckets at 4B, fits in a 64B cacheline. probably too aggressive
// even this should be practical enough for operator
var tiny = []uint32{
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

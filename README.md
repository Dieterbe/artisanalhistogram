## SimpleHistogram

this is an expirement. the reasoning/goals/tradeoffs need more validation.

### goals

* high performance histograms with a focus on understandability of the class intervals ("buckets"), eg have rounded intervals that show well on UI's
* optimize for typical range of networked applications, where we care for durations between roughly 1ms and 20s (and consider anything beyond equally bad)
* give equal weight to all samples within a given observation interval, and no weight to samples from prior intervals (contrast to EWMA based approaches)
* consistent bucket sizes so we can easily aggregate different histograms together (e.g. for timeseries rollups or runtime consolidation).
* don't be afraid to sacrifice accuracy to meet these goals. try to limit the error % but not to the extreme.

I *think* these requirements make https://github.com/codahale/hdrhistogram and https://github.com/VividCortex/gohistogram poor fits, but TBH I have to investigate them again.

First look at https://github.com/Dieterbe/simplehistogram/blob/master/cmd/play/main.go and then https://github.com/Dieterbe/simplehistogram/blob/master/hist2/hist2.go

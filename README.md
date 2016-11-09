## Artisanal histogram


Hand crafted histogram, made with love. To power insights from networked applications.  Not general-purpose.
Also somewhat experimental.


### goals

* reasonable performance (inserts now up to +- 20ns, could be improved more at cost of readability)
* optimize for typical range of networked applications, where we care for durations between roughly 1ms and 15s.
  anything under a ms is plenty fast.  Even if it was a a microsecond or less, we don't mind it being reported in the 1ms bucket.
  Likewise, anything over 10s is awful.  Whether it's 10s or 30s. Both are terrible and can go in the same bucket.
  Contrast this to [hdrhistograms](https://github.com/codahale/hdrhistogram) which are designed to provide buckets which can provide close approximations over huge ranges which I don't actually care about.
  This way we can also store the data in a more compact fashion.
* understandability of the class intervals ("buckets"), eg have rounded intervals that show well on UI's.
  powers of two are [faster to compute](http://pvk.ca/Blog/2015/06/27/linear-log-bucketing-fast-versatile-simple/) but then your buckets are like 1024, 1280, etc.
  I want to be able to answer questions like "how many requests were completed within 5 milliseconds? how many in a second or less"?
  Every histogram can return percentiles with a given degree of error.
  We allow for a bit more error in the typical case (and much more error for extreme outliers such as <<1ms and >>15s) in return for accurate numbers in histograms the way people actually want to look at them.
* consistent bucket sizes across different histograms so we can easily aggregate different histograms together (e.g. for timeseries rollups or runtime consolidation).
(this rules out [gohistogram](https://github.com/VividCortex/gohistogram)
* give equal weight to all samples within a given observation interval, and no weight to samples from prior intervals (contrast to EWMA based approaches)


### buckets

the following classes (buckets) have been manually crafted in an effort to
* cover the space well
* represent boundaries people actually care about
* still minimize errors as well as possible, by scaling up the class intervals corresponding to the bucket boundaries.

boundaries are in ms.
exactly 32 buckets. (32x4=128B size)

```
1
2
3
5
7.5
10
15
20
30
40
50
65
80
100
150
200
300
400
500
650
800
1000
1500
2000
3000
4000
5000
6500
8000
10000
15000
inf
```

### implementation notes

* 32 buckets because that should fit nicely on graphical UI's + also round size of 128B
* math.MaxUint32, i.e. 4294967295 should be a reasonable count limit. to save space.
  (it's easy to warn operators if they get too close to the limit and make them upgrade to bigger structures
  or shorter their reporting interval, and we can invalidate numbers if we detect an almost overflow)

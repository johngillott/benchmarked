# gobenchmark

## encoding/json

```bash
go test -benchmem -run=^$ github.com/johngillott/gobenchmark/encoding/json -bench BenchmarkMarshal -count=5 | tee bench.txt && benchstat bench.txt
```

```bash
ok      github.com/johngillott/gobenchmark/encoding/json        65.047s
name                                                   time/op
MarshalEncoder_map_string_string_jingo_sorted-8          5.89µs ± 2%
MarshalEncoder_map_string_string_jingo_unsorted-8        1.07µs ± 2%
MarshalEncoder_map_string_string_segmentio_sorted-8      16.4µs ± 2%
MarshalEncoder_map_string_string_segmentio_unsorted-8    8.61µs ± 0%
Marshal_map_string_string/std-8                          16.8µs ± 0%
Marshal_map_string_string/jsoniterator_sorted-8          12.5µs ± 5%
Marshal_map_string_string/jsoniterator_unsorted-8        3.24µs ± 1%
Marshal_map_string_string/segmentio_sorted-8             16.6µs ± 1%
Marshal_map_string_string/jettison_sorted-8              9.24µs ± 3%
Marshal_map_string_string/jettison_unsorted-8            4.34µs ± 2%

name                                                   speed
MarshalEncoder_map_string_string_jingo_sorted-8         175MB/s ± 2%
MarshalEncoder_map_string_string_jingo_unsorted-8       963MB/s ± 2%
MarshalEncoder_map_string_string_segmentio_sorted-8    63.1MB/s ± 2%
MarshalEncoder_map_string_string_segmentio_unsorted-8   120MB/s ± 0%
Marshal_map_string_string/std-8                        61.6MB/s ± 0%
Marshal_map_string_string/jsoniterator_sorted-8        82.8MB/s ± 5%
Marshal_map_string_string/jsoniterator_unsorted-8       318MB/s ± 1%
Marshal_map_string_string/segmentio_sorted-8           62.2MB/s ± 1%
Marshal_map_string_string/jettison_sorted-8             112MB/s ± 3%
Marshal_map_string_string/jettison_unsorted-8           238MB/s ± 2%

name                                                   alloc/op
MarshalEncoder_map_string_string_jingo_sorted-8           0.00B     
MarshalEncoder_map_string_string_jingo_unsorted-8         0.00B     
MarshalEncoder_map_string_string_segmentio_sorted-8      2.74kB ± 0%
MarshalEncoder_map_string_string_segmentio_unsorted-8    2.62kB ± 0%
Marshal_map_string_string/std-8                          5.70kB ± 0%
Marshal_map_string_string/jsoniterator_sorted-8          7.13kB ± 0%
Marshal_map_string_string/jsoniterator_unsorted-8        1.29kB ± 0%
Marshal_map_string_string/segmentio_sorted-8             3.89kB ± 0%
Marshal_map_string_string/jettison_sorted-8              1.15kB ± 0%
Marshal_map_string_string/jettison_unsorted-8            1.22kB ± 0%

name                                                   allocs/op
MarshalEncoder_map_string_string_jingo_sorted-8            0.00     
MarshalEncoder_map_string_string_jingo_unsorted-8          0.00     
MarshalEncoder_map_string_string_segmentio_sorted-8        91.0 ± 0%
MarshalEncoder_map_string_string_segmentio_unsorted-8      88.0 ± 0%
Marshal_map_string_string/std-8                            93.0 ± 0%
Marshal_map_string_string/jsoniterator_sorted-8            55.0 ± 0%
Marshal_map_string_string/jsoniterator_unsorted-8          4.00 ± 0%
Marshal_map_string_string/segmentio_sorted-8               92.0 ± 0%
Marshal_map_string_string/jettison_sorted-8                1.00 ± 0%
Marshal_map_string_string/jettison_unsorted-8              2.00 ± 0%
```

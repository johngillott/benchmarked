# gobenchmark

## encoding/json

```bash
go test -benchmem -run=^$ github.com/johngillott/gobenchmark/encoding/json -bench BenchmarkMarshal -count=5 | tee bench.txt && benchstat bench.txt
```

```bash
ok      github.com/johngillott/gobenchmark/encoding/json        65.758s
name                                                      time/op
MarshalEncoder_map_string_string_jingo_sorted-8             6.03µs ± 2%
MarshalEncoder_map_string_string_jingo_unsorted-8           1.09µs ± 2%
MarshalEncoder_map_string_string_segmentio_sorted-8         17.4µs ± 6%
MarshalEncoder_map_string_string_segmentio_unsorted-8       9.19µs ± 7%
MarshalEncoder_map_string_string/std-8                      17.9µs ± 6%
MarshalEncoder_map_string_string/jsoniterator_sorted-8      12.7µs ± 3%
MarshalEncoder_map_string_string/jsoniterator_unsorted-8    3.24µs ± 1%
MarshalEncoder_map_string_string/segmentio_sorted-8         17.7µs ±10%
MarshalEncoder_map_string_string/jettison_sorted-8          9.44µs ± 3%
MarshalEncoder_map_string_string/jettison_unsorted-8        4.52µs ± 4%

name                                                      speed
MarshalEncoder_map_string_string_jingo_sorted-8            171MB/s ± 2%
MarshalEncoder_map_string_string_jingo_unsorted-8          950MB/s ± 2%
MarshalEncoder_map_string_string_segmentio_sorted-8       59.5MB/s ± 6%
MarshalEncoder_map_string_string_segmentio_unsorted-8      113MB/s ± 7%
MarshalEncoder_map_string_string/std-8                    57.9MB/s ± 6%
MarshalEncoder_map_string_string/jsoniterator_sorted-8    81.2MB/s ± 3%
MarshalEncoder_map_string_string/jsoniterator_unsorted-8   318MB/s ± 1%
MarshalEncoder_map_string_string/segmentio_sorted-8       58.5MB/s ± 9%
MarshalEncoder_map_string_string/jettison_sorted-8         109MB/s ± 3%
MarshalEncoder_map_string_string/jettison_unsorted-8       228MB/s ± 4%

name                                                      alloc/op
MarshalEncoder_map_string_string_jingo_sorted-8              0.00B     
MarshalEncoder_map_string_string_jingo_unsorted-8            0.00B     
MarshalEncoder_map_string_string_segmentio_sorted-8         2.74kB ± 0%
MarshalEncoder_map_string_string_segmentio_unsorted-8       2.62kB ± 0%
MarshalEncoder_map_string_string/std-8                      5.70kB ± 0%
MarshalEncoder_map_string_string/jsoniterator_sorted-8      7.13kB ± 0%
MarshalEncoder_map_string_string/jsoniterator_unsorted-8    1.29kB ± 0%
MarshalEncoder_map_string_string/segmentio_sorted-8         3.89kB ± 0%
MarshalEncoder_map_string_string/jettison_sorted-8          1.15kB ± 0%
MarshalEncoder_map_string_string/jettison_unsorted-8        1.22kB ± 0%

name                                                      allocs/op
MarshalEncoder_map_string_string_jingo_sorted-8               0.00     
MarshalEncoder_map_string_string_jingo_unsorted-8             0.00     
MarshalEncoder_map_string_string_segmentio_sorted-8           91.0 ± 0%
MarshalEncoder_map_string_string_segmentio_unsorted-8         88.0 ± 0%
MarshalEncoder_map_string_string/std-8                        93.0 ± 0%
MarshalEncoder_map_string_string/jsoniterator_sorted-8        55.0 ± 0%
MarshalEncoder_map_string_string/jsoniterator_unsorted-8      4.00 ± 0%
MarshalEncoder_map_string_string/segmentio_sorted-8           92.0 ± 0%
MarshalEncoder_map_string_string/jettison_sorted-8            1.00 ± 0%
MarshalEncoder_map_string_string/jettison_unsorted-8          2.00 ± 0%
```

# gobenchmarked

## encoding/json

```bash
go test -benchmem -run=^$ github.com/johngillott/gobenchmarked/encoding/json -bench Benchmark -count=5 | tee bench.txt && benchstat bench.txt
```

```bash
ok      github.com/johngillott/gobenchmarked/encoding/json      67.153s
name                                               time/op
Marshal_map_string_string_jingo_sorted-8             5.93µs ± 3%
Marshal_map_string_string_jingo_unsorted-8           1.09µs ± 2%
Marshal_map_string_string_segmentio_sorted-8         16.4µs ± 1%
Marshal_map_string_string_segmentio_unsorted-8       8.68µs ± 1%
Marshal_map_string_string/std-8                      17.2µs ± 1%
Marshal_map_string_string/jsoniterator_sorted-8      12.4µs ± 1%
Marshal_map_string_string/jsoniterator_unsorted-8    3.33µs ± 4%
Marshal_map_string_string/segmentio_sorted-8         16.9µs ± 2%
Marshal_map_string_string/jettison_sorted-8          9.24µs ± 1%
Marshal_map_string_string/jettison_unsorted-8        4.36µs ± 1%

name                                               speed
Marshal_map_string_string_jingo_sorted-8            174MB/s ± 2%
Marshal_map_string_string_jingo_unsorted-8          947MB/s ± 2%
Marshal_map_string_string_segmentio_sorted-8       62.9MB/s ± 1%
Marshal_map_string_string_segmentio_unsorted-8      119MB/s ± 1%
Marshal_map_string_string/std-8                    60.2MB/s ± 1%
Marshal_map_string_string/jsoniterator_sorted-8    83.5MB/s ± 1%
Marshal_map_string_string/jsoniterator_unsorted-8   311MB/s ± 4%
Marshal_map_string_string/segmentio_sorted-8       61.0MB/s ± 2%
Marshal_map_string_string/jettison_sorted-8         112MB/s ± 1%
Marshal_map_string_string/jettison_unsorted-8       237MB/s ± 1%

name                                               alloc/op
Marshal_map_string_string_jingo_sorted-8              0.00B     
Marshal_map_string_string_jingo_unsorted-8            0.00B     
Marshal_map_string_string_segmentio_sorted-8         2.74kB ± 0%
Marshal_map_string_string_segmentio_unsorted-8       2.62kB ± 0%
Marshal_map_string_string/std-8                      5.70kB ± 0%
Marshal_map_string_string/jsoniterator_sorted-8      7.12kB ± 0%
Marshal_map_string_string/jsoniterator_unsorted-8    1.29kB ± 0%
Marshal_map_string_string/segmentio_sorted-8         3.89kB ± 0%
Marshal_map_string_string/jettison_sorted-8          1.15kB ± 0%
Marshal_map_string_string/jettison_unsorted-8        1.22kB ± 0%

name                                               allocs/op
Marshal_map_string_string_jingo_sorted-8               0.00     
Marshal_map_string_string_jingo_unsorted-8             0.00     
Marshal_map_string_string_segmentio_sorted-8           91.0 ± 0%
Marshal_map_string_string_segmentio_unsorted-8         88.0 ± 0%
Marshal_map_string_string/std-8                        93.0 ± 0%
Marshal_map_string_string/jsoniterator_sorted-8        55.0 ± 0%
Marshal_map_string_string/jsoniterator_unsorted-8      4.00 ± 0%
Marshal_map_string_string/segmentio_sorted-8           92.0 ± 0%
Marshal_map_string_string/jettison_sorted-8            1.00 ± 0%
Marshal_map_string_string/jettison_unsorted-8          2.00 ± 0%
```

# Golang Key/Value Engine Benchmark

**This is a work in progress, numbers are not final. We'd love to hear your feedback on how to optimize the code**

## Test

- key 16 bytes (like UUID)
- value 100 bytes (e.g. Email address + password)

## Test database:

- [flydb](https://github.com/ByteStorage/FlyDB)
- [bbolt](https://github.com/etcd-io/bbolt)
- [goleveldb](https://github.com/syndtr/goleveldb)
- [nutsdb](https://github.com/nutsdb/nutsdb)
- [rosedb](https://github.com/flower-corp/rosedb)
- [badger](https://github.com/dgraph-io/badger)
- [pebble](https://github.com/cockroachdb/pebble)
- [lotusdb](https://github.com/lotusdblabs/lotusdb)
- [pogreb](https://github.com/akrylysov/pogreb)
- [buntdb](https://github.com/tidwall/buntdb)

## Results

```
goos: linux / WSL
goarch: amd64
pkg: kvengines
cpu: AMD Ryzen 9 3900X 12-Core Processor
ssd: WD_BLACK SN850 1TB
Benchmark/Pogreb_1000000x_100b_nofsync_Put      1000000     4135 ns/op   1486 B/op 4 allocs/op
Benchmark/Pogreb_1000000x_100b_nofsync_Get      1000000      349 ns/op    112 B/op 1 allocs/op
Benchmark/Bbolt_1000000x_100b_nofsync_Put       1000000    27859 ns/op  24511 B/op        112 allocs/op
Benchmark/Bbolt_1000000x_100b_nofsync_Get       1000000     1746 ns/op    695 B/op23 allocs/op
Benchmark/BadgerDB_1000000x_100b_nofsync_Put    1000000     8604 ns/op   1879 B/op38 allocs/op
Benchmark/BadgerDB_1000000x_100b_nofsync_Get    1000000     5290 ns/op   1279 B/op20 allocs/op
Benchmark/BuntDB_1000000x_100b_nofsync_Put      1000000     6768 ns/op    876 B/op10 allocs/op
Benchmark/BuntDB_1000000x_100b_nofsync_Get      1000000     2408 ns/op    200 B/op 4 allocs/op
Benchmark/FlyDB_1000000x_100b_nofsync_Put       1000000     4136 ns/op    412 B/op 7 allocs/op
Benchmark/FlyDB_1000000x_100b_nofsync_Get       1000000     4441 ns/op    552 B/op 8 allocs/op
Benchmark/GoLevelDB_1000000x_100b_nofsync_Put   1000000     5800 ns/op    143 B/op 3 allocs/op
Benchmark/GoLevelDB_1000000x_100b_nofsync_Get   1000000     8437 ns/op   1073 B/op18 allocs/op
Benchmark/LotusDB_1000000x_100b_nofsync_Put     1000000    10624 ns/op   1913 B/op40 allocs/op
Benchmark/LotusDB_1000000x_100b_nofsync_Get     1000000     6297 ns/op    955 B/op26 allocs/op
Benchmark/NutsDB_1000000x_100b_nofsync_Put      1000000     5965 ns/op   1281 B/op17 allocs/op
Benchmark/NutsDB_1000000x_100b_nofsync_Get      1000000     6602 ns/op   1032 B/op15 allocs/op
Benchmark/Pebble_1000000x_100b_nofsync_Put      1000000     3194 ns/op    136 B/op 0 allocs/op
Benchmark/Pebble_1000000x_100b_nofsync_Get      1000000    13591 ns/op   3715 B/op 6 allocs/op
Benchmark/ROSEDB_1000000x_100b_nofsync_Put      1000000     4252 ns/op    156 B/op 5 allocs/op
Benchmark/ROSEDB_1000000x_100b_nofsync_Get      1000000     3893 ns/op    409 B/op 6 allocs/op
PASS
ok      kvengines       156.768s
```

## Note

Forked from [Contrast Benchmark](https://github.com/ByteStorage/contrast-benchmark)

Refactored towards a structure more like [Kvbench](https://github.com/smallnest/kvbench)

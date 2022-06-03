# Benchmark of various ordered structs in go.

## Legend

 - HashMap - not orderd, just an etalon for comparasion (HashMap)
 - RadixTree - github.com/gammazero/radixtree (Radix Tree with array nodes)
 - PlarRadixTree - github.com/plar/go-adaptive-radix-tree (Adaptive Radix Tree)
 - SortedSet - github.com/recoilme/sortedset (Array)
 - BTreeSet - github.com/recoilme/btreeset (Simple BTree)
 - Sorted - github.com/buraksezer/sorted (SkipList)
 - GoogleBTree - github.com/google/btree (LLRB Tree)
 - Art - github.com/arriqaaq/art (Adaptive Radix Tree)

## Results

 ```
(base) v-kulibaba:bench_sortedsets v.kulibaba$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/recoilme/bench_sortedsets
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkHashMapPut-8           26603949                62.44 ns/op            6 B/op          0 allocs/op
BenchmarkHashMapGet-8           66423618                32.66 ns/op            0 B/op          0 allocs/op
BenchmarkRadixTreePut-8         11432383               104.4 ns/op            44 B/op          1 allocs/op
BenchmarkRadixTreeGet-8         16536664                74.63 ns/op            0 B/op          0 allocs/op
BenchmarkPlarRadixTreePut-8      8326839               149.2 ns/op            36 B/op          1 allocs/op
BenchmarkPlarRadixTreeGet-8     11179678               121.5 ns/op             0 B/op          0 allocs/op
BenchmarkSortedSetPut-8          9485239               148.0 ns/op             4 B/op          0 allocs/op
BenchmarkSortedSetGet-8          8579707               149.7 ns/op             0 B/op          0 allocs/op
BenchmarkBTreeSetPut-8          11138948               125.1 ns/op             8 B/op          0 allocs/op
BenchmarkBTreeSetGet-8           9803560               139.2 ns/op             0 B/op          0 allocs/op
BenchmarkSortedPut-8             3925093               298.0 ns/op            32 B/op          0 allocs/op
BenchmarkSortedGet-8             3393968               367.9 ns/op             0 B/op          0 allocs/op
BenchmarkGoogleBTreePut-8        3482337               381.9 ns/op            33 B/op          1 allocs/op
BenchmarkGoogleBTreeGet-8        4549512               312.8 ns/op            16 B/op          1 allocs/op
BenchmarkArtPut-8               14644272                85.53 ns/op           25 B/op          1 allocs/op
BenchmarkArtGet-8               16931152                66.83 ns/op           15 B/op          0 allocs/op
PASS
 ```

# Conclusion

 github.com/arriqaaq/art - is best


# Author
 @recoilme
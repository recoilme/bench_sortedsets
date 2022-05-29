# Benchmark of various ordered structs in go.

## Legend

 - HashMap - not orderd, just an etalon for comparasion (HashMap)
 - RadixTree - github.com/gammazero/radixtree (Adaptive Radix Tree)
 - PlarRadixTree - github.com/plar/go-adaptive-radix-tree (Adaptive Radix Tree)
 - SortedSet - github.com/recoilme/sortedset (Array)
 - BTreeSet - github.com/recoilme/btreeset (Simple BTree)
 - Sorted - github.com/buraksezer/sorted (SkipList)
 - GoogleBTree - github.com/google/btree (LLRB Tree)

## Results

 ```
 (base) v-kulibaba:bench_sortedsets v.kulibaba$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/recoilme/bench_sortedsets
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkHashMapPut-8           26684005                61.04 ns/op            6 B/op          0 allocs/op
BenchmarkHashMapGet-8           60778569                32.14 ns/op            0 B/op          0 allocs/op
BenchmarkRadixTreePut-8         12444097                99.81 ns/op           43 B/op          1 allocs/op
BenchmarkRadixTreeGet-8         15230964                74.45 ns/op            0 B/op          0 allocs/op
BenchmarkPlarRadixTreePut-8      8196140               146.0 ns/op            36 B/op          1 allocs/op
BenchmarkPlarRadixTreeGet-8     10410999               122.7 ns/op             0 B/op          0 allocs/op
BenchmarkSortedSetPut-8          9648045               160.1 ns/op             4 B/op          0 allocs/op
BenchmarkSortedSetGet-8          8808595               145.4 ns/op             0 B/op          0 allocs/op
BenchmarkBTreeSetPut-8          12887174               110.7 ns/op             8 B/op          0 allocs/op
BenchmarkBTreeSetGet-8          10431980               129.8 ns/op             0 B/op          0 allocs/op
BenchmarkSortedPut-8             4385426               278.8 ns/op            33 B/op          0 allocs/op
BenchmarkSortedGet-8             4008255               354.0 ns/op             0 B/op          0 allocs/op
BenchmarkGoogleBTreePut-8        3718650               362.9 ns/op            32 B/op          1 allocs/op
BenchmarkGoogleBTreeGet-8        4640356               304.0 ns/op            16 B/op          1 allocs/op
PASS
 ```

# Conclusion

 github.com/gammazero/radixtree - best implementation by speed/allocations


# Author
 @recoilme
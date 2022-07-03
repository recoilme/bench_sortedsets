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
 - ArtRecoilme - github.com/recoilme/art (Adaptive Radix Tree)

## Results

 ```
(base) v-kulibaba:bench_sortedsets v.kulibaba$ go test -bench=. -benchmem -count=3 > x.txt
(base) v-kulibaba:bench_sortedsets v.kulibaba$ benchstat x.txt
name                time/op
HashMapPut-8         295ns ±18%
HashMapGet-8         111ns ±17%
RadixTreePut-8       200ns ± 1%
RadixTreeGet-8      78.6ns ± 1%
PlarRadixTreePut-8   304ns ± 3%
PlarRadixTreeGet-8   132ns ± 2%
SortedSetPut-8       243ns ± 2%
SortedSetGet-8       208ns ± 2%
BTreeSetPut-8        159ns ± 1%
BTreeSetGet-8        168ns ± 4%
SortedPut-8          386ns ± 4%
SortedGet-8          421ns ± 5%
GoogleBTreePut-8     643ns ± 7%
GoogleBTreeGet-8     320ns ± 4%
ArtPut-8             153ns ± 0%
ArtGet-8            47.9ns ±10%
ArtRecoilmePut-8     124ns ± 3%
ArtRecoilmeGet-8    34.0ns ± 3%

name                alloc/op
HashMapPut-8         68.3B ±13%
HashMapGet-8         0.00B     
RadixTreePut-8        129B ± 0%
RadixTreeGet-8       0.00B     
PlarRadixTreePut-8    107B ± 0%
PlarRadixTreeGet-8   0.00B     
SortedSetPut-8       35.0B ± 0%
SortedSetGet-8       0.00B     
BTreeSetPut-8        64.0B ± 0%
BTreeSetGet-8        0.00B     
SortedPut-8           212B ± 4%
SortedGet-8          0.00B     
GoogleBTreePut-8      151B ± 0%
GoogleBTreeGet-8     16.0B ± 0%
ArtPut-8             83.0B ± 0%
ArtGet-8             0.00B     
ArtRecoilmePut-8     90.0B ± 0%
ArtRecoilmeGet-8     0.00B     

name                allocs/op
HashMapPut-8          0.00     
HashMapGet-8          0.00     
RadixTreePut-8        2.00 ± 0%
RadixTreeGet-8        0.00     
PlarRadixTreePut-8    4.00 ± 0%
PlarRadixTreeGet-8    0.00     
SortedSetPut-8        0.00     
SortedSetGet-8        0.00     
BTreeSetPut-8         0.00     
BTreeSetGet-8         0.00     
SortedPut-8           0.00     
SortedGet-8           0.00     
GoogleBTreePut-8      3.00 ± 0%
GoogleBTreeGet-8      1.00 ± 0%
ArtPut-8              3.00 ± 0%
ArtGet-8              0.00     
ArtRecoilmePut-8      1.00 ± 0%
ArtRecoilmeGet-8      0.00  
 ```



# Author
 @recoilme
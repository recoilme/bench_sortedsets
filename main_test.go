package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
	"time"

	art2 "github.com/arriqaaq/art"
	"github.com/buraksezer/sorted"
	"github.com/gammazero/radixtree"
	"github.com/google/btree"
	art "github.com/plar/go-adaptive-radix-tree"
	art4 "github.com/recoilme/art"

	"github.com/recoilme/btreeset"
	"github.com/recoilme/sortedset"
)

func initStr(N int) []string {
	strs := make([]string, N)
	for n := 0; n < N; n++ {
		strs[n] = strconv.FormatInt(int64(n), 10)
	}
	return strs
}

func initBin(N int) [][]byte {
	bytes := make([][]byte, N)
	for n := 0; n < N; n++ {
		bytes[n] = make([]byte, 8)
		binary.BigEndian.PutUint64(bytes[n], uint64(n))
	}
	return bytes
}

func BenchmarkHashMapPut(b *testing.B) {
	strs := initStr(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	set := make(map[string]bool)
	for n := 0; n < b.N; n++ {
		set[strs[n]] = true
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	strs := initStr(b.N)
	set := make(map[string]bool)
	for n := 0; n < b.N; n++ {
		set[strs[n]] = true
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = set[strs[n]]
	}
}

func BenchmarkRadixTreePut(b *testing.B) {
	strs := initStr(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	tree := radixtree.New()
	for n := 0; n < b.N; n++ {
		tree.Put(strs[n], nil)
	}
}

func BenchmarkRadixTreeGet(b *testing.B) {
	strs := initStr(b.N)

	tree := radixtree.New()
	for n := 0; n < b.N; n++ {
		tree.Put(strs[n], nil)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, found := tree.Get(strs[n])
		if found != true {
			fmt.Println("Error, not found:", strs[n])
		}
	}
}

func BenchmarkPlarRadixTreePut(b *testing.B) {
	strs := initStr(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	tree := art.New()
	for n := 0; n < b.N; n++ {
		tree.Insert(art.Key(strs[n]), nil)
	}
}

func BenchmarkPlarRadixTreeGet(b *testing.B) {
	strs := initStr(b.N)
	tree := art.New()
	for n := 0; n < b.N; n++ {
		tree.Insert(art.Key(strs[n]), nil)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, found := tree.Search(art.Key(strs[n]))
		if found != true {
			fmt.Println("Error, not found:", strs[n])
		}
	}
}

func BenchmarkSortedSetPut(b *testing.B) {
	strs := initStr(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	set := sortedset.New()
	for n := 0; n < b.N; n++ {
		set.Put(strs[n])
	}
}

func BenchmarkSortedSetGet(b *testing.B) {
	strs := initStr(b.N)
	set := sortedset.New()
	for n := 0; n < b.N; n++ {
		set.Put(strs[n])
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		found := set.Has(strs[n])
		if found != true {
			fmt.Println("Error, not found:", strs[n])
		}
	}
}

func BenchmarkBTreeSetPut(b *testing.B) {
	bytes := initBin(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	bt := &btreeset.BTreeSet{}
	for n := 0; n < b.N; n++ {
		bt.Set(bytes[n])
	}
}

func BenchmarkBTreeSetGet(b *testing.B) {
	bytes := initBin(b.N)

	bt := &btreeset.BTreeSet{}
	for n := 0; n < b.N; n++ {
		bt.Set(bytes[n])
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		found := bt.Has(bytes[n])
		if found != true {
			fmt.Println("Error, not found:", string(bytes[n]))
		}
	}
}

func BenchmarkSortedPut(b *testing.B) {
	bytes := initBin(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	ss := sorted.NewSortedSet(0)
	//defer ss.Close()
	for n := 0; n < b.N; n++ {
		ss.Set(bytes[n])
	}
}

func BenchmarkSortedGet(b *testing.B) {
	bytes := initBin(b.N)

	ss := sorted.NewSortedSet(0)
	//defer ss.Close()
	for n := 0; n < b.N; n++ {
		ss.Set(bytes[n])
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		found := ss.Check(bytes[n]) //bt.Has(bytes[n])
		if found != true {
			fmt.Println("Error, not found:", string(bytes[n]))
		}
	}
}

// Google need an interface
// Str implements the Item interface for strings.
type Str string

// Less returns true if a < b.
func (a Str) Less(b btree.Item) bool {
	return a < b.(Str)
}

func BenchmarkGoogleBTreePut(b *testing.B) {
	strs := initStr(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	tr := btree.New(3)
	for n := 0; n < b.N; n++ {
		tr.ReplaceOrInsert(Str(strs[n]))
	}
}

func BenchmarkGoogleBTreeGet(b *testing.B) {
	strs := initStr(b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	tr := btree.New(3)
	for n := 0; n < b.N; n++ {
		tr.ReplaceOrInsert(Str(strs[n]))
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {

		found := tr.Has(Str(strs[n]))
		if found != true {
			fmt.Println("Error, not found:", string(strs[n]))
		}
	}
}

func BenchmarkArtPut(b *testing.B) {

	strs := initBin(b.N)

	b.ResetTimer()
	b.ReportAllocs()
	tree := art2.NewTree()
	for n := 0; n < b.N; n++ {
		tree.Insert(strs[n], nil)
	}
}

func BenchmarkArtGet(b *testing.B) {
	strs := initBin(b.N)

	tree := art2.NewTree()
	for n := 0; n < b.N; n++ {
		tree.Insert(strs[n], strs[n])
	}

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		val := tree.Search(strs[n])
		if !bytes.Equal(val.([]byte), strs[n]) {
			b.Fail()
		}

	}
}

func BenchmarkArtRecoilmePut(b *testing.B) {

	strs := initBin(b.N)

	b.ResetTimer()
	b.ReportAllocs()
	tree := art4.New()
	for n := 0; n < b.N; n++ {
		tree.Set(strs[n], nil)
	}
}

func BenchmarkArtRecoilmeGet(b *testing.B) {
	strs := initBin(b.N)

	tree := art4.New()
	for n := 0; n < b.N; n++ {
		tree.Set(strs[n], nil)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = tree.Get(strs[n])

	}
}

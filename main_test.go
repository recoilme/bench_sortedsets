package main

import (
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
	"github.com/recoilme/btreeset"
	"github.com/recoilme/sortedset"
)

func BenchmarkHashMapPut(b *testing.B) {
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	b.ResetTimer()
	b.ReportAllocs()
	set := make(map[string]bool)
	for n := 0; n < b.N; n++ {
		set[strs[n]] = true
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
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
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	b.ResetTimer()
	b.ReportAllocs()
	tree := radixtree.New()
	for n := 0; n < b.N; n++ {
		tree.Put(strs[n], nil)
	}
}

func BenchmarkRadixTreeGet(b *testing.B) {
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

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
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	b.ResetTimer()
	b.ReportAllocs()
	tree := art.New()
	for n := 0; n < b.N; n++ {
		tree.Insert(art.Key(strs[n]), nil)
	}
}

func BenchmarkPlarRadixTreeGet(b *testing.B) {
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

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
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	b.ResetTimer()
	b.ReportAllocs()
	set := sortedset.New()
	for n := 0; n < b.N; n++ {
		set.Put(strs[n])
	}
}

func BenchmarkSortedSetGet(b *testing.B) {
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

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
	bytes := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		bytes[n] = []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
	b.ResetTimer()
	b.ReportAllocs()
	bt := &btreeset.BTreeSet{}
	for n := 0; n < b.N; n++ {
		bt.Set(bytes[n])
	}
}

func BenchmarkBTreeSetGet(b *testing.B) {
	bytes := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		bytes[n] = []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	}

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
	bytes := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		bytes[n] = []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
	b.ResetTimer()
	b.ReportAllocs()
	ss := sorted.NewSortedSet(0)
	//defer ss.Close()
	for n := 0; n < b.N; n++ {
		ss.Set(bytes[n])
	}
}

func BenchmarkSortedGet(b *testing.B) {
	bytes := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		bytes[n] = []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	}

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
	strs := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		strs[n] = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	b.ResetTimer()
	b.ReportAllocs()
	tr := btree.New(3)
	for n := 0; n < b.N; n++ {
		tr.ReplaceOrInsert(Str(strs[n]))
	}
}

func BenchmarkGoogleBTreeGet(b *testing.B) {
	strs := make([]string, b.N)
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
	
	strs := make([][]byte, b.N)
	
	for n := 0; n < b.N; n++ {
		bin:=make([]byte,8)
		binary.BigEndian.PutUint64(bin,uint64(time.Now().UnixNano()))
		strs[n] = bin
	}

	b.ResetTimer()
	b.ReportAllocs()
	tree := art2.NewTree()
	for n := 0; n < b.N; n++ {
		tree.Insert(strs[n], nil)
	}
}

func BenchmarkArtGet(b *testing.B) {
	strs := make([][]byte, b.N)
	
	for n := 0; n < b.N; n++ {
		bin:=make([]byte,8)
		binary.BigEndian.PutUint64(bin,uint64(time.Now().UnixNano()))
		strs[n] = bin
	}

	tree := art2.NewTree()
	for n := 0; n < b.N; n++ {
		tree.Insert(strs[n], 0)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = tree.Search(strs[n])
		
	}
}

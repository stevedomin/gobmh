package gobmh

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// --- Tests

type IndexTestCase struct {
	h, n  string
	index int
}

var testCases = []IndexTestCase{
	{"", "", 0},
	{"f", "foo", -1},
	{"foo", "f", 0},
	{"bar", "f", -1},
	{"foofoofoogo", "go", 9},
	{"foobafoofoobarfoobarbar", "foobar", 8},
}

func TestIndexHorspool(t *testing.T) {
	testAllCases(t, IndexHorspool, testCases)
}

func testAllCases(t *testing.T, f func(h, n []byte) int, testCases []IndexTestCase) {
	for _, test := range testCases {
		h := []byte(test.h)
		n := []byte(test.n)
		result := f(h, n)
		if result != test.index {
			t.Errorf("f(%q, %q) = %v. Expect : %v", h, n, result, test.index)
		}
	}
}

// --- Benchmarks

var haystack []byte

type BenchCase struct {
	n     string
	index int
}

var benchCases = []BenchCase{
	{"INTRODUCTION TO MATHEMATICAL PHILOSOPHY", 24},
	{"religion", 13863},
	{"was a very voluminous writer, mainly", 868248},
	{"actions are good which in fact promote the general happiness", 1915558},
	{"Zeno the Ekatic", 2202542},
}

func init() {
	var err error

	haystack, err = ioutil.ReadFile("westernphilosophy.txt")
	if err != nil {
		panic("Test file can't be read")
	}
}

func BenchmarkIndex(b *testing.B) {
	benchmarkCases(b, bytes.Index, benchCases)
}

func BenchmarkIndexHorspool(b *testing.B) {
	benchmarkCases(b, IndexHorspool, benchCases)
}

func benchmarkCases(b *testing.B, f func(h, n []byte) int, benchCases []BenchCase) {
	var n []byte
	var result, i int

	for _, bench := range benchCases {
		n = []byte(bench.n)
		result = f(haystack, n)
		if result != bench.index {
			panic("Index function is not working as expected.")
		}

		for i = 0; i < b.N; i++ {
			_ = f(haystack, n)
		}
	}
}

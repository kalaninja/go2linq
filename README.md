# go2linq [![GoDoc](https://godoc.org/github.com/kalaninja/go2linq?status.svg)](https://godoc.org/github.com/kalaninja/go2linq) [![Build Status](https://travis-ci.org/kalaninja/go2linq.svg?branch=master)](https://travis-ci.org/kalaninja/go2linq) [![Coverage Status](https://coveralls.io/repos/github/kalaninja/go2linq/badge.svg?branch=master)](https://coveralls.io/github/kalaninja/go2linq?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kalaninja/go2linq)](https://goreportcard.com/report/github.com/kalaninja/go2linq)
A powerful language integrated query (LINQ) library for Go.
* Safe for concurrent use
* Complete lazy evaluation
* Supports arrays, slices, maps, strings, channels and
custom collections (collection needs to implement Iterable interface
and element - Comparable interface)
* Parallel LINQ (PLINQ) *(comming soon)*

## Installation

    $ go get github.com/kalaninja/go2linq

## Quickstart

Usage is as easy as chaining methods like

`From(slice)` `.Where(predicate)` `.Select(selector)` `.Union(data)` 

Just keep writing.

**Example:** Find all owners of cars manufactured from 2015
```go
import . "github.com/kalaninja/go2linq"
	
type Car struct {
    id, year int
    owner, model string
}

owners := []string{}

From(cars).Where(func(c interface{}) bool {
	return c.(Car).year >= 2015
}).Select(func(c interface{}) interface{} {
	return c.(Car).owner
}).ToSlice(&owners)
```

**Example:** Find an author that has written the most books
```go
import . "github.com/kalaninja/go2linq"
	
type Book struct {
	id      int
	title   string
	authors []string
}

author := From(books).SelectMany(func(b interface{}) Query {
		return From(b.(Book).authors)
	}).GroupBy(func(a interface{}) interface{} {
		return a
	}, func(a interface{}) interface{} {
		return a
	}).OrderByDescending(func(g interface{}) interface{} {
		return len(g.(Group).Group)
	}).Select(func(g interface{}) interface{} {
		return g.(Group).Key
	}).First()
```

**More examples** can be found in [documentation](https://godoc.org/github.com/kalaninja/go2linq)

## Performance

Due to lazy execution go2linq has better performance and allocates less memory
in most cases. Benchmark comparing go2linq to
[go-linq](https://github.com/ahmetalpbalkan/go-linq) is available in
[benchmark_test.go](https://github.com/kalaninja/go2linq/blob/master/benchmark_test.go).
Below is the result of this benchmark on my machine (MacBookPro8,1 Intel Core i5 2,4 GHz):
```
BenchmarkSelectWhereFirst-4       	 3000000	       561 ns/op	     224 B/op	      10 allocs/op
BenchmarkSelectWhereFirst_golinq-4	       2	 555810859 ns/op	120546360 B/op	 2000085 allocs/op
BenchmarkSum-4                    	      20	  73847428 ns/op	 8000289 B/op	 1000019 allocs/op
BenchmarkSum_golinq-4             	       5	 253731714 ns/op	69161392 B/op	 1000053 allocs/op
BenchmarkZipSkipTake-4            	 5000000	       351 ns/op	     192 B/op	       6 allocs/op
BenchmarkZipSkipTake_golinq-4     	       2	 672403213 ns/op	144520824 B/op	 3000075 allocs/op
```
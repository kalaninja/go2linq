package go2linq

import (
	"testing"

	golinq "github.com/ahmetalpbalkan/go-linq"
)

const (
	size = 1000000
)

func BenchmarkSelectWhereFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, size).Select(func(i interface{}) interface{} {
			return -i.(int)
		}).Where(func(i interface{}) bool {
			return i.(int) > -1000
		}).First()
	}
}

func BenchmarkSelectWhereFirst_golinq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		golinq.Range(1, size).Select(func(a golinq.T) (golinq.T, error) {
			return -a.(int), nil
		}).Where(func(a golinq.T) (bool, error) {
			return a.(int) > -1000, nil
		}).First()
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, size).Where(func(i interface{}) bool {
			return i.(int)%2 == 0
		}).SumInts()
	}
}

func BenchmarkSum_golinq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		golinq.Range(1, size).Where(func(a golinq.T) (bool, error) {
			return a.(int)%2 == 0, nil
		}).Sum()
	}
}

func BenchmarkZipSkipTake(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, size).Zip(Range(1, size).Select(func(i interface{}) interface{} {
			return i.(int) * 2
		}), func(i, j interface{}) interface{} {
			return i.(int) + j.(int)
		}).Skip(2).Take(5)
	}
}

func BenchmarkZipSkipTake_golinq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		golinq.Range(1, size).Zip(golinq.Range(11, size).Select(func(i golinq.T) (golinq.T, error) {
			return i.(int) * 2, nil
		}), func(i, j golinq.T) (golinq.T, error) {
			return i.(int) + j.(int), nil
		}).Skip(2).Take(5)
	}
}

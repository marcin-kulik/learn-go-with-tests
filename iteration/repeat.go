package iteration

import "testing"

func Repeat(times int, letter string) string{
	var repeated string
	for i := 0; i< times; i++ {
		repeated += letter
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(3,"a")
	}
}
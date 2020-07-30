// Package benchmark_slice_append prvides benchmark tests for slice mofications.
package benchmark_slice_append

// Usage:
// go test benchmark_slice_append_test.go --bench . -benchmem

import (
  "testing"
)

var (
	c, ca []int
)

// BenchmarkSliceAppend benchmark test to append elements to initially nil slice.
// Delays caused when slice capacity is full and requires creating new slice
// with double capacity and copy to new slice.
func BenchmarkSliceAppend(b *testing.B) {
	v := make([]int, 10000000, 10000000)
	b.ResetTimer()
  for i := 0; i < b.N; i++ {
		c = nil
		for _, val := range v {
			c = append(c, val)
		}
  }
}

// BenchmarkSliceAppendPreAllocate benchmark test when no capacity modifications
// happens during copy and no new slice allocations and copy happens.
func BenchmarkSliceAppendPreAllocate(b *testing.B) {
	v := make([]int, 10000000, 10000000)
	b.ResetTimer()
  for i := 0; i < b.N; i++ {
		b.StopTimer()
		c := make([]int, 0, 10000000)
		b.StartTimer()
		for _, val := range v {
			c = append(c, val)
		}
  }
}

// BenchmarkSliceCopy benchmark test to copy slice when destination slice is
// pre-allocated.
func BenchmarkSliceCopy(b *testing.B) {
	v := make([]int, 10000000, 10000000)
	b.ResetTimer()
  for i := 0; i < b.N; i++ {
		c = make([]int, 10000000, 10000000)
		copy(c, v)
  }
}

// BenchmarkSliceCopyAppend benchmark test to copy slice when no new
// pre-allocation happens, but compiler trick used.
func BenchmarkSliceCopyAppend(b *testing.B) {
	v := make([]int, 10000000, 10000000)
	b.ResetTimer()
  for i := 0; i < b.N; i++ {
		ca = append([]int(nil), v...)
  }
}

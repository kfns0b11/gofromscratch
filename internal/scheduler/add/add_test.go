package add

import (
	"math/rand"
	"testing"
	"time"
)

func generateNums(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(n)
	}
	return nums
}

func BenchmarkAdd(b *testing.B) {
	nums := generateNums(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(nums)
	}
}

func BenchmarkAddConcurrency(b *testing.B) {
	nums := generateNums(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addConcurency(8, nums)
	}
}

package add

import (
	"sync"
	"sync/atomic"
)

func add(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func addConcurency(gnum int, nums []int) int {
	var sum int64
	stride := len(nums) / gnum
	var wg sync.WaitGroup
	wg.Add(gnum)

	for i := 0; i < gnum; i++ {
		go func(i int) {
			defer wg.Done()
			var tmpsum int
			start := i * stride
			end := start + stride
			if i == gnum-1 {
				end = len(nums)
			}
			for _, v := range nums[start:end] {
				tmpsum += v
			}

			atomic.AddInt64(&sum, int64(tmpsum))
		}(i)
	}
	wg.Wait()
	return int(sum)
}

package concurrency

import (
	"sync"
)

func syncWithMutexAdd() int {
	var count int
	var wg sync.WaitGroup
	wg.Add(8)
	var mu sync.Mutex

	for i := 0; i < 8; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return count
}

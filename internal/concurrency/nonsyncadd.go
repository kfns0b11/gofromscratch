package concurrency

import (
	"sync"
)

func nonSyncAdd() int {
	var count int
	var wg sync.WaitGroup
	wg.Add(8)

	for i := 0; i < 8; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	return count
}

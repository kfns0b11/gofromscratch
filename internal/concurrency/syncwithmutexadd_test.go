package concurrency

import "testing"

func TestSyncWithMutexAdd(t *testing.T) {
	count := syncWithMutexAdd()
	if count != 8*10000 {
		t.Errorf("Expected count to be 80000, but got %d", count)
	}
}

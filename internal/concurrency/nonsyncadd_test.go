package concurrency

import "testing"

func TestNonSyncAdd(t *testing.T) {
	count := nonSyncAdd()
	if count != 8*10000 {
		t.Errorf("Expected count to be 80000, but got %d", count)
	}
}

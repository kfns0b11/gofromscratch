// singleton lazy instantiation, means that global singleton
// instance is created when it first used. It is non-concurrently
// secure and requires lock in practice
package lazysingleton

import "sync"

type singleton struct {
}

var ins *singleton
var mu sync.Mutex

func GetInsOr() *singleton {
	if ins == nil {
		mu.Lock()
		if ins == nil {
			return &singleton{}
		}
		mu.Unlock()
	}
	return ins
}

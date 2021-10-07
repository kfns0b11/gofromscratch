// Implementate singleton design pattern elegantly in go with
// the help of sync.Once
package elegantsingleton

import "sync"

type singleton struct {
}

var ins *singleton
var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

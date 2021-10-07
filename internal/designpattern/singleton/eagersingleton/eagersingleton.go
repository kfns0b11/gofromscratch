// singleton eager instantiation, means that global singleton
// instance is created when the package is loaded. It will
// cause longer to load if singleton type is time-consuming.
package eagersingleton

type singleton struct {
}

var ins *singleton = &singleton{}

func GetInsOr() *singleton {
	return ins
}

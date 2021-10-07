package methodfactory

import (
	"fmt"
	"testing"
)

func TestNewPersonFactory(t *testing.T) {
	newbaby := NewPersonFactory(1)
	baby := newbaby("lily")
	fmt.Printf("%+v", baby)
}

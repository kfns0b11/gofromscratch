package options

import (
	"fmt"
	"testing"
)

func TestOptions(t *testing.T) {
	wc := WithCaching(true)
	wt := WithTimeout(10)

	con, err := NewConnect("localhost", wc, wt)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", con)
	}
}

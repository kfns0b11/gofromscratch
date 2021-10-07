package strategy

import (
	"fmt"
	"testing"
)

func TestOperator(t *testing.T) {
	operator := Operator{}
	operator.SetStrategy(&Add{})
	res := operator.Caculate(1, 2)
	fmt.Println(res)

	operator.SetStrategy(&Reduce{})
	res = operator.Caculate(1, 2)
	fmt.Println(res)
}

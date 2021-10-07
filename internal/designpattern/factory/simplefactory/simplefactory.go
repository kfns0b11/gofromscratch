// Simple factory return a struct instance.
// The simple factory pattern ensures that the instances we
// create have the required parameters, and thus the methods
// of the instances can be executed as expected.
package simplefactory

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi, my name is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

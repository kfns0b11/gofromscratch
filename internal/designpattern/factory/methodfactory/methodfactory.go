// Method factory return a func to construct the struct instance
package methodfactory

type Person struct {
	Name string
	Age  int
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			Name: name,
			Age:  age,
		}
	}
}

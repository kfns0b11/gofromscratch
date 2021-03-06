// Abstract factory pattern.
// The only difference between it and the simple factory pattern
// is that it returns an interface instead of a structure.
package abstractfactory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi, my name is %s", p.name)
}

func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}

// ----------------------------------------------------------

// We define a Doer interface, that has the method signature
// of the `http.Client` structs `Do` method
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// This gives us a regular HTTP client from the `net/http` package
func NewHTTPClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// The `NewRecorder` method of the httptest package gives us
	// a new mock request generator
	res := httptest.NewRecorder()

	// calling the `Result` method gives us
	// the default empty *http.Response object
	return res.Result(), nil
}

// This gives us a mock HTTP client, which returns
// an empty response for any request sent to it
func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}

func QueryUser(doer Doer) error {
	req, err := http.NewRequest("Get", "http://iam.api.marmotedu.com:8080/v1/secrets", nil)
	if err != nil {
		return err
	}

	_, err = doer.Do(req)
	if err != nil {
		return err
	}

	return nil
}

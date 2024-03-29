package gosamplelib

import (
	"fmt"
)

// Person struct
type Person struct {
	Name string
	Age int
}

// Say function
func (p *Person) Say() {
	fmt.Printf("My name is %s, I'm %d\n", p.Name, p.Age)
}

// Say2 function
func (p *Person) Say2() (message string) {
	message = fmt.Sprintf("My name is %s, I'm %d", p.Name, p.Age)
	return
}

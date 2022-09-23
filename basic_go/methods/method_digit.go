package methods

import (
	"fmt"
	"strings"
)

// source: https://www.digitalocean.com/community/tutorials/defining-methods-in-go

// Defining a Method
/*
	In other languages, the receiver of method invocations is typically referred to by a keyword (e.g. this or self)

	We created a struct called "Creature" with string fields for a "Name" and a "Greeting".
	This "Creature" has a single method defined, "Greet". Within the receiver declaration,
	we assigned the instance of "Creature" to the variable "c"
*/

type Creatrue struct {
	Name     string
	Greeting string
}

func (c Creatrue) Greet() Creatrue {
	fmt.Printf("%s says %s\n", c.Name, c.Greeting)
	return c
}

func (c Creatrue) SayGoodbye(name string) {
	fmt.Println("Farewell", name, "!\n")
}

/*
Both styles output the same results, but the example using dot notation is far
more readable. The chain of dots also tells us the sequence in which methods
will be invoked, where the functional style inverts this sequence.
*/
func Execute_define_method_digitoc() {
	sammy := Creatrue{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	fmt.Printf("%p\n", &sammy)
	m := sammy.Greet() // ~= Creatrue.Greet(sammy)
	fmt.Printf("%p\n", &m)
	sammy.Greet().SayGoodbye("gophers")
	Creatrue.SayGoodbye(Creatrue.Greet(sammy), "gophers")
}

// Interface
/*
	When you define a method on any type in Go, that method is added to the "type’s method set".

	The method set is the collection of functions associated with that type as methods
	and used by the Go compiler to determine whether some type can be assigned to
	a variable with an interface type.

	An interface type is a specification of methods used by the compiler to guarantee
	that a type provides implementations for those methods.

	Any type that has methods with the same name, same parameters, and same return
	values as those found in an interface’s definition are said to implement that
	interface and are allowed to be assigned to variables with that interface’s type.

	so far, we have defined methods on the value receiver. That is, if we use the
	functional invocation of methods, the first parameter, referring to the type the
	method was defined on, will be a value of that type, rather than a pointer.
	Consequently, any modifications we make to the instance provided to the method
	will be discarded when the method completes execution, because the value
	received is a copy of the data.
*/
type Stringer interface {
	String() string
}

type Ocean struct {
	Creatures []string
}

func (o Ocean) String() string {
	return strings.Join(o.Creatures, ", ")
}

func log(header string, s fmt.Stringer) {
	fmt.Println(header, ":", s)
}

func Execute_interface_digitoc() {
	/*
		"Ocean" is said to implement the "fmt.Stringer" interface
	*/
	o := Ocean{
		Creatures: []string{
			"sea urchin",
			"lobster",
			"shark",
		},
	}
	log("ocean contains", o)
}

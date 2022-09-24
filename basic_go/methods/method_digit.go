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

// Pointer Receivers
/*
	In this example, We want to force code in other packages to only add occupants with the AddOccupant method,
	, so we’ve made the occupants field unexported by lowercasing the first letter of the field name.

	We also want to make sure that calling AddOccupant will cause the instance of Boat
	to be modified, which is why we defined AddOccupant on the pointer receiver.

	Pointers act as a reference to a specific instance of a type rather than a copy of that type.
*/
type Boat struct {
	Name      string
	occupants []string
}

func (b *Boat) AddOccupant(name string) *Boat {
	b.occupants = append(b.occupants, name)
	return b
}

func (b Boat) Manifest() {
	fmt.Println("The", b.Name, "has the following occupants:")
	for _, n := range b.occupants {
		fmt.Println("\t", n)
	}
}

func Execute_Pointer_receiver() {
	b := &Boat{
		Name: "S.S. DigitalOcean",
	}
	fmt.Printf("%p\n", b)
	b.AddOccupant("sammy the Shark")
	b.AddOccupant("Larry the Lobster")

	/*
		The "Manifest" method is defined on the "Boat" value, because in its definition,
		the receiver is specified as "(b Boat)".

		we are still able to call Manifest because Go is able to automatically dereference the pointer
		to obtain the "Boat" value.

		"b.Manifest()" here is equivalent to "(*b).Manifest()".
	*/
	b.Manifest()
}

// Pointer Receivers and Interface
/*
	The method sets for the pointer receiver and the value receiver are different
	because methods that receive a pointer can modify their receiver where those
	that receive a value cannot.

	We defined a Dive() method on the pointer receiver to Shark which modified
	isUnderwater to true.

	We also defined the "String()" method of the value receiver so that it could
	cleanly print the state of the "Shark" using "fmt.Println"
	by using the "fmt.Stringer" interface accepted by "fmt.Println".

	Using the Submersible interface rather than a *Shark allows the submerge function to
	depend only on the behavior provided by a type.
*/

type Submersible interface {
	Dive()
}

type Shark struct {
	Name string

	isUnderwater bool
}

func (s Shark) String() string {
	if s.isUnderwater {
		return fmt.Sprintf("%s is underwater", s.Name)
	}
	return fmt.Sprintf("%s is on the surface", s.Name)
}

func (s *Shark) Dive() {
	s.isUnderwater = true
}

func submerge(s Submersible) {
	s.Dive()
}

func Execute_pointer_receivers_interface() {
	s := &Shark{
		Name: "Sammy",
	}
	fmt.Println(s)

	submerge(s)

	fmt.Println(s)
}

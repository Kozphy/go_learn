package _interface

import (
	"fmt"
	"math"
)

func All_interface() {
	fmt.Println("Interface===================")
	/*
		An interface type is defined as a set of method signatures.

		A value of interface type can hold any value that implements those methods.
	*/
	var abser Abser
	var f_a = MyFloat(-math.Sqrt2)
	var v5 = Vertex{X: 3, Y: 4}

	abser = f_a // a MyFloat implements Abser
	abser = &v5 // a *Vertex implements Abser

	// In the following line, v5 is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// abser = v5
	fmt.Println(abser.Abs(), "\n")

	fmt.Println("Interface are implemented implicitly")
	// A type implements an interface by implementing its methods.
	var i I = &T{S: "hello"}
	i.M()

	fmt.Println("\n")
	fmt.Println("Interface values")
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
	// (value, type)
	var i_p I = &T{S: "Hello"}
	Describe(i_p)
	i_p.M()

	i = F(math.Pi)
	Describe(i)
	i.M()
	fmt.Println("\n")
	fmt.Println("Interface values with nil underlying values")
	var i2 I
	var t2_p *T

	i2 = t2_p
	Describe(i2)
	i2.M()

	i2 = &T{S: "hello"}
	Describe(i2)
	i2.M()
	fmt.Println()
	fmt.Println("Nil interface values")
	var i3 I
	Describe(i3)
	// Calling a method on a nil interface is a run-time error because
	// there is no type inside the interface tuple to indicate which concrete method to call.
	// i3.M()
	fmt.Println()
	fmt.Println("The empty interface")
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	var i4 interface{}
	describe(i4)

	i4 = 42
	describe(i4)

	i4 = "hello"
	describe(i4)
	fmt.Println()
	fmt.Println("Type assertions")
	var i5 interface{} = "hello"
	/*
		A type assertion provides access to an interface value's underlying concrete value.
		t := i.(T)
		This statement asserts that the interface value i holds the concrete type T and
		assigns the underlying T value to the variable t.

		If i holds a T, then t will be the underlying value and ok will be true.

		If not, ok will be false and t will be the zero value of type T, and no panic occurs.

		If i does not hold a T and doesn't declare ok syntax, the statement will trigger a panic.
	*/
	s5 := i5.(string)
	fmt.Println(s5)

	s5, ok := i5.(string)
	fmt.Println(s5, ok)

	f5, ok := i5.(float64)
	fmt.Println(f5, ok)

	// f5 = i5.(float64) // panic
	// fmt.Println(f5)
	fmt.Println()
	fmt.Println("Type switches")
	/*
		A type switch is like a regular switch statement, but the cases in a type switch specify
		types (not values), and those values are compared against the type of the value held
		by the given interface value.
	*/
	do(21)
	do("hello")
	do(true)

	fmt.Println()
	fmt.Println("Stringers")
	/*
		type Stringer interface {
			String() string
		}
		A Stringer is a type that can describe itself as a string.
		The fmt package (and many others) look for this interface to print values.
	*/
	a := Person{Name: "Arthur Dent", Age: 42}
	z := Person{Name: "Zaphod Beeblebrox", Age: 9001}
	fmt.Println(a, z)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	/*
		The declaration in a type switch has the same syntax as a type assertion "i.(T)",
		but the specific type "T" is replaced with the keyword "type".

		the variable v is of the same interface type and value as i.
	*/
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func Describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

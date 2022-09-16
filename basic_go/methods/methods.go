package methods

import (
	"fmt"
	"math"
)

func All_methods() {
	fmt.Println("methods==============================")
	fmt.Println("Methods")
	/*
		A method is a function with a special "receiver" argument.
		The "receiver" appears in its own argument list between the "func" keyword and the method name.
		You can only declare a method with a receiver whose type is defined in the same package as the method.
		You cannot declare a method with a receiver whose type is defined in another package
	*/
	var v = Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())

	fl := MyFloat(-math.Sqrt2)
	fmt.Println(fl.Abs(), "\n")

	fmt.Println("Pointer receivers")
	v.Scale(10)
	fmt.Println(v.Abs(), "\n")

	fmt.Println("Methods and pointer indirection")
	var v2 = Vertex{X: 3, Y: 4}
	// Go interprets the statement v2.Scale(2) as (&v2).Scale(2) since the Scale method has a pointer receiver.
	v2.Scale(2)
	ScaleFunc(&v2, 10)
	var v2_p = &Vertex{X: 4, Y: 3}
	v2_p.Scale(3)
	ScaleFunc(v2_p, 8)
	fmt.Println(v2, v2_p, "\n")

	fmt.Println("Methods and pointer indirection (2)")
	var v3 = Vertex{X: 3, Y: 4}
	fmt.Println(v3.Abs())
	fmt.Println(AbsFunc(v3), "\n")
	var v3_p = &Vertex{X: 3, Y: 4}
	// In this case, the method call v3_p.Abs() is interpreted as (*v3_p).Abs().
	fmt.Println(v3_p.Abs())
	fmt.Println(AbsFunc(*v3_p), "\n")

	fmt.Println("Choosing a value or pointer receiver")
	var v4_p = &Vertex{X: 3, Y: 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v4_p, v4_p.Abs_p())
	v4_p.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v4_p, v4_p.Abs_p())

	fmt.Println("\n")
}

type Abser interface {
	Abs_p() float64
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Abs_p() float64 {
	/*
		There are two reasons to use a pointer receiver.
		The first is so that the method can modify the value that its receiver points to.

		The second is to avoid copying the value on each method call.
		This can be more efficient if the receiver is a large struct, for example.
	*/
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	/*
		Methods with pointer receivers can modify the value to which the receiver points

		With a value receiver, the Scale method operates on a copy of the original Vertex value.
		(This is the same behavior as for any other function argument.)
		The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
	*/

	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f *MyFloat) Abs_p() float64 {
	if (*f) < 0 {
		return float64(-(*f))
	}
	return float64(-(*f))
}

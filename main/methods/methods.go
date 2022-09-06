package methods

import (
	"math"
)

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

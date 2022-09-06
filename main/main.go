package main

import (
	"fmt"
	"log"
	"math"

	"example.com/hello"
	functions "github.com/zixas/go_learn/func"
	"github.com/zixas/go_learn/greetings"
	_interface "github.com/zixas/go_learn/interface"
	mapping "github.com/zixas/go_learn/map"
	"github.com/zixas/go_learn/methods"
	pointers "github.com/zixas/go_learn/pointer"
	"github.com/zixas/go_learn/slice"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")

	fmt.Println("==============================")
	fmt.Println(quote.Go())

	fmt.Println("==============================")
	hello.Hello()

	names := []string{"Gladys", "Samantha", "Darrin"}
	var message, err = greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("==============================")
	fmt.Println(message)
	fmt.Println("Map==============================")
	fmt.Println("Map intro")
	mapping.Map_intro()
	fmt.Println("Pointer==============================")
	fmt.Println("P_basic")
	pointers.P_basic()

	fmt.Println("Pointers to struct")
	pointers.Pointers_to_struct()

	fmt.Println("Struct literals")
	pointers.Struct_Literals()
	fmt.Println("Slice==============================")
	fmt.Println("slice literals")
	slice.Slice_literals()

	fmt.Println("Slice of slices")
	slice.Slice_of_slices()
	fmt.Println("Function==============================")
	fmt.Println("Function values")
	functions.Function_values()

	fmt.Println("function closures")
	functions.Function_closures()

	fmt.Println("function fiboncci")
	f := functions.Fiboncci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println("methods==============================")
	fmt.Println("Methods")
	/*
		A method is a function with a special "receiver" argument.
		The "receiver" appears in its own argument list between the "func" keyword and the method name.
		You can only declare a method with a receiver whose type is defined in the same package as the method.
		You cannot declare a method with a receiver whose type is defined in another package
	*/
	var v = methods.Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())

	fl := methods.MyFloat(-math.Sqrt2)
	fmt.Println(fl.Abs(), "\n")

	fmt.Println("Pointer receivers")
	v.Scale(10)
	fmt.Println(v.Abs(), "\n")

	fmt.Println("Methods and pointer indirection")
	var v2 = methods.Vertex{X: 3, Y: 4}
	// Go interprets the statement v2.Scale(2) as (&v2).Scale(2) since the Scale method has a pointer receiver.
	v2.Scale(2)
	methods.ScaleFunc(&v2, 10)
	var v2_p = &methods.Vertex{X: 4, Y: 3}
	v2_p.Scale(3)
	methods.ScaleFunc(v2_p, 8)
	fmt.Println(v2, v2_p, "\n")

	fmt.Println("Methods and pointer indirection (2)")
	var v3 = methods.Vertex{X: 3, Y: 4}
	fmt.Println(v3.Abs())
	fmt.Println(methods.AbsFunc(v3), "\n")
	var v3_p = &methods.Vertex{X: 3, Y: 4}
	// In this case, the method call v3_p.Abs() is interpreted as (*v3_p).Abs().
	fmt.Println(v3_p.Abs())
	fmt.Println(methods.AbsFunc(*v3_p), "\n")

	fmt.Println("Choosing a value or pointer receiver")
	var v4_p = &methods.Vertex{X: 3, Y: 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v4_p, v4_p.Abs_p())
	v4_p.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v4_p, v4_p.Abs_p())

	fmt.Println("\n")
	fmt.Println("Interface")
	/*
		An interface type is defined as a set of method signatures.

		A value of interface type can hold any value that implements those methods.
	*/
	var abser _interface.Abser
	var f_a = _interface.MyFloat(-math.Sqrt2)
	var v5 = _interface.Vertex{X: 3, Y: 4}

	abser = f_a // a MyFloat implements Abser
	abser = &v5 // a *Vertex implements Abser

	// In the following line, v5 is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// abser = v5
	fmt.Println(abser.Abs(), "\n")

	fmt.Println("Interface are implemented implicitly")
	// A type implements an interface by implementing its methods.
	var i _interface.I = _interface.T{"hello"}
	i.M()

	fmt.Println("\n")
	fmt.Println("Interface values")

}

package reflection

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func Law_reflect() {
	fmt.Println("Reflection")
	Repre_interface()
	Reflect_first_law()
	Reflect_second_law()
	Reflect_third_law()
	Struct_with_reflect()

}

// article source: https://go.dev/blog/laws-of-reflection

/*

	Reflection in computing is the ability of a program to examine its own structure, particularly through types;
	it’s a form of metaprogramming.

*/

// Types and interfaces
/*
	If we declare following, then  i has type int and j has type MyInt.
	The variables i and j have distinct static types and,
	although they have the same underlying type,
	they cannot be assigned to one another without a conversion.

	An interface variable can store any concrete (non-interface) value as long as
	that value implements the interface’s methods.

*/
func Type_and_interface() {
	type MyInt int

	var i int
	var y MyInt
	_ = i
	_ = y

	// It's important to be clear that whatever concrete value r may hold,
	// r's type is always io.Reader:
	// Go is statically typed and the static type of r is io.Reader
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
}

// The representation of and interface
/*
	A variable of interface type stores a pair: the concrete value assigned to the variable,
	and that value’s type descriptor.

	To be more precise, the value is the underlying concrete data item that
	implements the interface and the type describes the full type of that item.

*/
/*
	r contains, schematically, the (value, type) pair, (tty, *os.File).
	Notice that the type *os.File implements methods other than Read;
	even though the interface value provides access only to the Read method,
	the value inside carries all the type information about that value.
*/

func Repre_interface() (interface{}, error) {
	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	r = tty
	describe(r)

	var w io.Writer
	w = r.(io.Writer)
	describe(w)
	var empty interface{}
	empty = w
	_ = empty
	describe(empty)
	return r, err
}

func describe(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

// Reflection goes from interface value to reflection object
/*
	At the basic level, reflection is just a mechanism to examine the type and value pair stored inside an interface variable

	Those two types (type, value) give access to the contents of an interface variable,
	and two simple functions, called reflect.TypeOf and reflect.ValueOf, retrieve reflect.Type
	and reflect.Value pieces out of an interface value.
*/
func Reflect_first_law() {
	fmt.Println()
	fmt.Println("Reflect first law: ")
	var x float64 = 3.4
	/*
		When we call reflect.TypeOf(x), x is first stored in an empty interface, which is then passed as the argument;
		reflect.TypeOf unpacks that empty interface to recover the type information.
	*/
	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())

	/*
		the “getter” and “setter” methods of Value operate on the largest
		type that can hold the value: int64 for all the signed integers.

		for instance. That is, the Int method of Value returns an int64 and
		the SetInt value takes an int64;

	*/
	var y uint8 = 'x'
	v1 := reflect.ValueOf(y)
	fmt.Println("type: ", v1.Type())
	fmt.Println("kind is uint8: ", v1.Kind() == reflect.Uint8)
	y = uint8(v1.Uint()) // v1.Uint returns a uint64
	fmt.Println()

	/*
		Kind of a reflection object describes the underlying type, not the static type.
		If a reflection object contains a value of a user-defined integer type, as in

		the "Kind" of v is still reflect.Int, even though the static type of x is MyInt, not int. In other words,
		the "Kind" cannot discriminate an int from a MyInt even though the Type can.
	*/
	type MyInt int
	var x1 MyInt = 7
	v2 := reflect.ValueOf(x1)
	fmt.Println(v2.String())
	fmt.Println(v2.Type())
	fmt.Println(v2.Kind())
}

// Reflection goes from reflection object to interface value.
func Reflect_second_law() {
	fmt.Println()
	fmt.Println("Reflect second law: ")
	/*
		Given a reflect.Value we can recover an interface value using the Interface method
		in effect the method packs the type and value information
		back into an interface representation and returns the result

		there’s no need to type-assert the result of v.Interface() to float64;
		the empty interface value has the concrete value’s type information inside and Printf will recover it.

		In short, the Interface method is the inverse of the ValueOf function,
		except that its result is always of static type interface{}.
	*/
	var x float64 = 10.1
	v := reflect.ValueOf(x)
	y := v.Interface().(float64) // y will have type float64
	fmt.Println(y)
	fmt.Println(v.Interface())
	fmt.Printf("value is %e\n", v.Interface())
}

// To modify a reflection object, the value must be settable.
func Reflect_third_law() {
	fmt.Println()
	fmt.Println("Reflect third law: ")
	var x float64 = 3.4
	// we pass  a copy of x to reflect.ValueOf, so the interface value created
	// as the argument to reflect.ValueOf is a copy of x, not x itself.
	v := reflect.ValueOf(x)
	/*
		The problem is not that the value 7.1 is not addressable;
		it’s that v is not settable. Settability is a property of a reflection Value,
		and not all reflection Values have it
	*/
	// v.SetFloat(7.1) // will panic

	// The "CanSet" method of Value reports the settability of a Value; in out case,
	fmt.Println("settability of v: ", v.CanSet())

	/*
		Settability is a bit like addressability, but stricter.
		It’s the property that a reflection object can modify the actual storage
		that was used to create the reflection object.
		Settability is determined by whether the reflection object holds the original item.

		If we want to modify x by reflection, we must give the reflection library a pointer
		to the value we want to modify.

		The reflection object p isn't settable, but it's not p we want to set, it's *p.
		To get to what p points to, we call the "Elem" method of Value, which indirects through the pointer,
		and save the result in a reflection Value called v.

		keep in mind that reflection Values need the address of something in order to modify what they represent.
	*/
	p := reflect.ValueOf(&x)
	fmt.Println("type of p: ", p.Type())
	fmt.Println("settability of p: ", p.CanSet())
	v = p.Elem()
	fmt.Println("settability of v: ", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v)
	fmt.Println(v.Interface())
	fmt.Println(x)

}

// Struct with reflect
func Struct_with_reflect() {
	fmt.Println()
	fmt.Println("Struct with reflect")
	// There’s one more point about settability introduced in passing here:
	// the field names of T are upper case (exported) because only exported
	// fields of a struct are settable.
	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/hello"
	generic "github.com/zixas/go_learn/Generics"
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
	var i _interface.I = &_interface.T{S: "hello"}
	i.M()

	fmt.Println("\n")
	fmt.Println("Interface values")
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
	// (value, type)
	var i_p _interface.I = &_interface.T{S: "Hello"}
	_interface.Describe(i_p)
	i_p.M()

	i = _interface.F(math.Pi)
	_interface.Describe(i)
	i.M()
	fmt.Println("\n")
	fmt.Println("Interface values with nil underlying values")
	var i2 _interface.I
	var t2_p *_interface.T

	i2 = t2_p
	_interface.Describe(i2)
	i2.M()

	i2 = &_interface.T{S: "hello"}
	_interface.Describe(i2)
	i2.M()
	fmt.Println()
	fmt.Println("Nil interface values")
	var i3 _interface.I
	_interface.Describe(i3)
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
	a := _interface.Person{Name: "Arthur Dent", Age: 42}
	z := _interface.Person{Name: "Zaphod Beeblebrox", Age: 9001}
	fmt.Println(a, z)

	fmt.Println("Exercise: Stringers")

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println()
	fmt.Println("Errors")
	/*
		"error" type is a built-in interface similar to fmt.Stringer
		type error interface {
			Error() string
		}
	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("Exercise: Errors")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	fmt.Println()
	fmt.Println("Reader")
	r := strings.NewReader("Hello, Reader")

	b := make([]byte, 8)
	for {
		/*
			type Reader interface {
				Read(p []byte) (n int, err error)
			}
			Read reads up to len(p) bytes into p.
			It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
			Even if Read returns n < len(p), it may use all of p as scratch space during the call

			When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read and
			Reader may return either err == EOF or err == nil. The next Read should return 0, EOF.
		*/
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
	fmt.Println("Exercise: rot13 Reader")
	s := strings.NewReader("Lbh penpxrq gur phqr!")
	rot := &rot13Reader{s}
	/*
		If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst).
		Otherwise, if dst implements the ReaderFrom interface, the copy is implemented by calling dst.ReadFrom(src).
	*/
	io.Copy(os.Stdout, rot)

	fmt.Println()
	fmt.Println()
	fmt.Println("Images")
	im := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(im.Bounds())
	fmt.Println(im.At(0, 0).RGBA())

	fmt.Println()
	fmt.Println("Generics=====================")
	fmt.Println("Type parameters")
	/*
		The type parameters of a function appear between brackets, before the function's arguments.
		func Index[T comparable](s []T, x T) int

		"comparable" is a useful constraint that makes it possible to use the == and != operators on values of the type.
	*/
	si := []int{10, 20, 15, -10}
	fmt.Println(generic.Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(generic.Index(ss, "hello"))

	fmt.Println()
	fmt.Println("Generic types")
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i, c := range p {
		switch {
		case c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm':
			p[i] += 13
		case c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z':
			p[i] -= 13
		}
	}
	return
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

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var res = make([]string, len(ip))
	for i, add := range ip {
		res[i] = strconv.Itoa(int(add))
	}
	return strings.Join(res, ".")
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil

}

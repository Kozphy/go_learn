package pointer

import "fmt"

// source: https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go

// Introduction
/*
	You pass data to function is call argument.

	Sometimes, the function needs a local copy of the data, and you want the original
	to remain unchanged.

	This is called passing by value, because you’re sending the value of the variable to the function, but not the variable itself.

	Other times, you may want the function to be able to alter the data in the original variable.

	In this case, you don’t need to send the actual data to the function; you just need to tell the function where the data is located in memory.

	This is called passing by reference, because the value of the variable isn’t passed to the function, just its location.

	A data type called a pointer holds the memory address of the data, but not the data itself.
*/

// Defining and Using Pointers
/*
	If you place an ampersand(&) in front of a variable name, you are stating that
	you want to get the address, or a pointer to that variable.

	The second syntax element is the use of the asterisk (*) or dereferencing operator.

	The memory address is a hexadecimal number, and not meant to be human-readable.
*/

func Execute_define_using_pointer() {
	var creature string = "shark"
	// pointer variable is storing the "address" of the "creature" variable.
	var pointer *string = &creature
	fmt.Println("creature =", creature)
	// we received the address of where the creature variable is currently stored in computer memory.
	fmt.Println("address of pointer variable value =", pointer)
	// we received the address of where pointer variable is currently stored in computer memory.
	fmt.Printf("address of pointer variable = %p\n", &pointer)

	// dereference
	fmt.Println("*pointer = ", *pointer)

	// change value through deference
	*pointer = "jellyfish"
	fmt.Println("*pointer = ", *pointer)

	fmt.Println("creature = ", creature)
}

// Function Pointer Receivers
type Creature struct {
	Species string
}

func Execute_function_pointer_receivers() {
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature_with_no_check(&creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature_with_no_check(creature *Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

func changeCreature_with_check(creature *Creature) {
	// This is a common approach for checking for nil:
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

// Nil Pointers
/*
	All variables in Go have a zero value. This is true even for a pointer.
	If you declare a pointer to a type, but assign no value, the zero value will be
	nil. nil is a way to say that “nothing has been initialized” for the variable.

	When we run the program, it printed out the value of the creature variable, and the value is <nil>.

	it panics. This is because there is no instance of the variable actually created. Because of this,
	the program has no where to actually store the value, so the program panics.
*/
func Execute_nil_pointers() {
	var creature *Creature
	fmt.Printf("1) %+v\n", creature)
	changeCreature_with_no_check(creature)
	fmt.Printf("3) %+v\n", creature)
}

/*
When you are working with pointers, there is a potential for the program to panic.
To avoid panicking, you should check to see if a pointer value is nil prior to trying
to access any of the fields or methods defined on it.
*/
func Execute_nil_pointer_improve() {
	var creature *Creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	changeCreature_with_check(creature)
	fmt.Printf("3) %+v\n", creature)

}

// Method Pointer Receivers
/*
	if you define a method with a value receiver, you are not able to make changes
	to the instance of that type that the method was defined on.

	If we want to be able to modify the instance of the "creature" variable in the methods,
	we need to define them as having a "pointer" receiver:
*/
func (c Creature) String() string {
	return c.Species
}

func (c *Creature) Reset() {
	c.Species = ""
}

func Execute_method_pointer_receivers() {
	var creature Creature = Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	creature.Reset()
	fmt.Printf("2) %+v\n", creature)
}

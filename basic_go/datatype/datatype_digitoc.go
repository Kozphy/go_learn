package datatype

// source: https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go

// Background
/*
	One way to think about data types is to consider the different types of data
	that we use in the real world. An example of data in the real world are
	numbers: we may use whole numbers (0, 1, 2, …), integers (…, -1, 0, 1, …), and
	irrational numbers (π).

	For computers, each data type is quite different—like words and numbers.
	As a result we have to be careful about how we use varying data types to assign
	values and how we manipulate them through operations.
*/

// Integers

// Floating-Point Numbers
/*
	A floating-point number or a float is used to represent "real numbers" that cannot
	be expressed as integers.

	Real numbers include all rational and irrational numbers, because of this, floating-point
	numbers can contain a fractional part, such as 9.0 or -116.42.


	source of real number: https://en.wikipedia.org/wiki/Real_number
*/

// Sizes of Numberic Types
/*
	Go has two types of numeric data that are distinguished by the static or dynamic
	nature of their sizes.

	The first type is an architecture-independent type, which means that the size of
	the data in bits does not change, regardless of the machine that the code is
	running on.

	Most system architectures today are either 32 bit or 64 bit. For instance, you
	may be developing for a modern Windows laptop, on which the operating system
	runs on a 64-bit architecture. However, if you are developing for a device like
	a fitness watch, you may be working with a 32-bit architecture. If you use an
	architecture-independent type like int32, regardless of the architecture you
	compile for, the type will have a constant size.

	The second type is an implementation-specific type. In this type, the bit size
	can vary based on the architecture the program is built on. For instance, if we
	use the int type, when Go compiles for a 32-bit architecture, the size of the
	data type will be 32 bits. If the program is compiled for a 64-bit architecture, the variable will be 64 bits in size.

	Go has the following architecture-independent integer types:
	uint8       unsigned  8-bit integers (0 to 255)
	uint16      unsigned 16-bit integers (0 to 65535)
	uint32      unsigned 32-bit integers (0 to 4294967295)
	uint64      unsigned 64-bit integers (0 to 18446744073709551615)
	int8        signed  8-bit integers (-128 to 127)
	int16       signed 16-bit integers (-32768 to 32767)
	int32       signed 32-bit integers (-2147483648 to 2147483647)
	int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	Floats and complex numbers also come in varying sizes:
	float32     IEEE-754 32-bit floating-point numbers
	float64     IEEE-754 64-bit floating-point numbers
	complex64   complex numbers with float32 real and imaginary parts
	complex128  complex numbers with float64 real and imaginary parts

	There are also a couple of alias number types, which assign useful names to specific data types:
	byte        alias for uint8
	rune        alias for int32

	The purpose of the byte alias is to make it clear when your program is using
	bytes as a common computing measurement in character string elements

	Even though byte and uint8 are identical once the program is compiled,
	byte is often used to represent character data in numeric form, whereas
	uint8 is intended to be a number in your program.

	The "rune" alias is a bit different. Where byte and uint8 are exactly the same
	data, a rune can be a single byte or four bytes, a range determined by int32.
	A rune is used to represent a Unicode character, whereas only ASCII characters
	can be represented solely by an int32 data type.

	In addition, Go has the following implementation-specific types:
	uint     unsigned, either 32 or 64 bits
	int      signed, either 32 or 64 bits
	uintptr  unsigned integer large enough to store the uninterpreted bits of a pointer value
*/

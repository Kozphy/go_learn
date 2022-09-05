package slice

import (
	"io/ioutil"
	"regexp"
)

// source article: https://go.dev/blog/slices-intro

// Array
/*
	The slice type is an abstraction built on top of Go’s array type,
	the type [4]int represents an array of four integers.
	An array’s size is fixed; its length is part of its type
	([4]int and [5]int are distinct, incompatible types).

	Arrays do not need to be initialized explicitly;
	the zero value of an array is a ready-to-use array whose elements are themselves zeroed

	when you assign or pass around an array value you will make a copy of its contents.
*/

// Slice
/*
	The type specification for a slice is []T, where T is the type of the elements of the slice.
	Unlike an array type, a slice type has no specified length.

	A slice literal is declared just like an array literal, except you leave out the element count.

	A slice can be created with the built-in function called make, which has the signature,
	func make([]T, len, cap) []T

	The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.

	A slice can also be formed by “slicing” an existing slice or array.
	Slicing is done by specifying a "half-open range [ : )" with two indices separated by a colon.
	ex: b[1:4]

	A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment,
	and its capacity (the maximum length of the segment).

	The length is the number of elements referred to by the slice.
	The capacity is the number of elements in the underlying array
	(beginning at the element referred to by the slice pointer)

	Slicing does not copy the slice’s data. It creates a new slice value that points to the original array.

	A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a slice or array.
*/

// similar func copy(dst, src []T) int
// To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.
func grow_slice() {
	s := make([]byte, 5, 5)
	t := make([]byte, len(s), (cap(s)+1)*2)
	for i := range s {
		t[i] = s[i]
	}
	s = t
}

// similar func append(s []T, x ...T) []T
// To append one slice to another, use ... to expand the second argument to a list of arguments.
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

// ===========================================
var digitRegexp = regexp.MustCompile("[0-9]+")

// possible gotcha
// re-slicing a slice doesn’t make a copy of the underlying array.
// so can't release memory until it is no longer referenced.
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

// improve version
// copy to new slice
func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

# dsa algo with go chp2

> last updated: 2022/11/01

## intro

In data structures, a collection of elements of a single type is called an **array**.

**Slices** are  similar to arrays except that they have unusual properties. Slice operations such as enlarging a slice using `append` and `copy` methods, assigning parts of a slice, appending a slice, and appending a part of a slice are presented with code samples.

In this chapter, we will discuss the following Go language-specific data structures:

- Arrays
- Slices
- Two-dimensional slices
- Maps
- Database operations
- Variadic functions
- CRUD web forms

## Array

Go arrays are not dynamic but have a `fixed size`. `To add more elements than the size`, a `bigger array needs to be created` and all the elements of the `old one need to be copied`.

An array is passed as a value through functions by `copying the array`, `not having` any provision of `inbuilt methods to increase it's size in Go`. Passing a big array to a function might be a `performance issue`.

## Slice

**Go Slice** is an abstraction over **Go Array**.

 A Go slice can be appended to elements after the capacity has reached its size.

 Slices are `dynamic` and can double the current capacity in order to add more elements

### len() function

The `len()` function gives the `current length of slice`, and the `capacity of slice` can `be obtained using the cap()` function.

### Slice function

Slices are `passed by referring` to functions. Big slices can be passed to functions `without impacting performance`.

```go
//twiceValue method given slice of int type
func twiceValue(slice []int) {

	var i int
	var value int

	for i, value = range slice {

		slice[i] = 2 * value

	}

}

// main method
func main() {

	var slice = []int{1, 3, 5, 6}
	twiceValue(slice)

	var i int

	for i = 0; i < len(slice); i++ {

		fmt.Println("new slice value", slice[i])
	}
}
```

## Two-dimensional slices

Two-dimensional slices are descriptors of a two-dimensional array.

A two-dimensional slice is a contiguous section of an array that is `stored away from the slice itself`. It `holds references to an underlying array`.

```go
var TwoDArray [8][8]int
TwoDArray[3][6] = 18
```

### two slice

```go
var rows int = 7
var cols int = 9
var twoslices = make([][]int, rows)
for i := range twoslices {
	twoslices[i] = make([]int, cols)
}
	fmt.Println(twoslices)
```

### append method

The `append` method on the slice is used to append new elements to the slice. If the slice `capacity has reached the size of the underlying array`, then append `increases the size by creating a new underlying array` and adding the new element.

```go
var arr = [] int{5,6,7,8,9}
var slic1 = arr[: 3]
fmt.Println("slice1",slic1)
var slic2 = arr[1:5]
fmt.Println("slice2",slic2)
var slic3 = append(slic2, 12)
fmt.Println("slice3",slic3)
```

## Maps

```go
var languages = map[int]string {
	3: "English",
	4: "Franch",
	5: "Spanish"
}

var products = make(map[int]string)
products[1] = "chair"
products[2] = "table"

for i, v := range languages {
	fmt.Println("language", i, ":", value)
}
fmt.Println("product with key 2", products[2])

fmt.Println(products[2])
delete(products, "chair")
fmt.Println("products", products)
```

## Variadic functions

A function in which we pass an infinite number of arguments, instead of passing them one at a time, is called a `variadic function`.

The type of the final parameter is preceded by an `ellipsis (...)`, while declaring a variadic function

```go

```

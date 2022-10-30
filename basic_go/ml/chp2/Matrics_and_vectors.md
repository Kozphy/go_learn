# Matrices, Probability, and Statistics

> last updated: 2022/10/30

## Matrices and vectors

For the most part, we will utilize packages from `github.com/gonum` to form and work with matrices and vectors

### Vectors

A vector is an ordered collection of numbers arranged in either a row (left to right) or column (up and down). Each of the numbers in a vector is called a component.

```go
var myvector []float64

myvector = append(myvector, 11.0)
myvector = append(myvector, 5.2)

fmt.Println(myvector)
```

Slices are indeed ordered collections. However, they don't really represent the concept of rows or columns, and we would still need to work out various vector operations on top of slices.

On the vector operation side, gonum provides `gonum.org/v1/gonum/floats` to operate on slices of float64 values and `gonum.org/v1/gonum/mat`, which ,along with, matrices provides a `Vector` type (with corresponding methods):

```go
// Create a new vector value.
myvector := mat.NewVector(2, []float64{11.0, 5.2})
```

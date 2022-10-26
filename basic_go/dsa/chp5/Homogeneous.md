# Homogeneous Data structure

## intro

**Homogeneous** data structures contain similar types of data, such as integers or double values. Homogeneous data structures are used in matrices, as well as tensor and vector mathematics.

**Tensors** are mathematical structures for **scalars** and **vectors**.

A first-rank tensor is a **vector**. A vector consists of a row or a column.

A tensor with zero rank is a **scalar**.

A **matrix** is a two-dimensional cluster of numbers.

## Two-dimensional arrays

Every element in a two-dimensional array arr, is identified as `arr[i][j]`, where arr is the name of the array and i and j `represent rows and columns`, and their values ranging from 0 to m and 0 to n, respectively. Traversing a two-dimensional array is of `O(m*n) complexity`.

```go
var arr = [4][5] int {
    {4,5,7,8,9},
    {1,2,4,5,6},
    {9,10,11,12,14},
    {3,5,6,8,9}
}
```

`Homogeneous` data structure `arrays` consist of `contiguous memory address locations`

`Two-dimensional matrices are modeled as two-dimensional arrays`.

A scalar is an element of a field that defines a vector space. A matrix can be multiplied by a scalar. You can divide a matrix by any non-zero real number.

The order of a matrix is the number of rows, m, by the number of columns, n.

A matrix with `rows m and columns n` is referred to `as` an **m x n matrix**.

There are multiple types of matrices, such as a **row matrix, column matrix, triangular matrix, null matrix, and zero matrix**; let's discuss them in the following sections.

## Row matrix

A **row matrix is a `1 x m` matrix** consisting of a single row of m elements, as shown here.

```go
var matrix = [1][3] int {
    {1,2,3}
}
```

## Column matrix

A **column matrix is an `m x 1` matrix** that has a single column of m elements.

```go
var matrix = [4][1] int {
    {1},
    {2},
    {3},
    {4}
}
```

## Lower triangular matrix

A **lower triangular matrix** consists of elements that have a `value of zero above` the **main diagonal**.

```go
var matrix = [3][3] int {
    {1,0,0},
    {1,1,0},
    {2,1,1}
}
```

## Upper triangular matrix

An **upper triangular matrix** consists of elements with a `value of zero below` the **main diagonal**.

```go
var matrix = [3][3] int {
    {1,2,3},
    {0,1,4},
    {0,0,1}
}
```

## Null matrix

A **null or a zero matrix** is a matrix `entirely consisting of zero values`, as shown in the following code snippet:

```go
var matrix = [3][3] int {
    {0,0,0},
    {0,0,0},
    {0,0,0}
}
```

## Identity matrix

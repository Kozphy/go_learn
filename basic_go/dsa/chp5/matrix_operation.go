package chp5

import "fmt"

func Execute_Matrix() {
	matrix1 := [2][2]int{
		{4, 5},
		{1, 2},
	}

	matrix2 := [2][2]int{
		{6, 7},
		{3, 4},
	}

	fmt.Printf("origin matrix1:%v , matrix2: %v\n", matrix1, matrix2)
	sum_matrix := add(matrix1, matrix2)
	fmt.Printf("sum matrix: %v\n", sum_matrix)

	substract_matrix := substract(matrix1, matrix2)
	fmt.Printf("substract matrix: %v\n", substract_matrix)

	multiply_matrix := multiply(matrix1, matrix2)
	fmt.Printf("multiply matrix: %v\n", multiply_matrix)

	transpose_matrix := transpose(matrix1)
	fmt.Printf("transpose matrix: %v\n", transpose_matrix)
}

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var sum [2][2]int

	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

func substract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var difference [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference
}

func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var product [2][2]int

	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			var productSum int = 0
			for n := 0; n < 2; n++ {
				productSum = productSum + matrix1[l][n]*matrix2[n][m]
			}
			product[l][m] = productSum
		}
	}
	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var transMatrix [2][2]int

	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

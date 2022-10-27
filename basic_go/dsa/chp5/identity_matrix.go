package chp5

import "fmt"

func Identity(order int) {
	var matrix [][]float64 = make([][]float64, order)

	for i := 0; i < order; i++ {
		var temp []float64 = make([]float64, order)
		temp[i] = 1
		matrix[i] = temp
		fmt.Println(matrix[i])
	}

}

func Execute_Identity() {
	Identity(4)
}

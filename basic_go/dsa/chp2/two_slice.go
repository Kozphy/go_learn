package chp2

import "fmt"

func two_slice() {
	var rows int = 7
	var cols int = 9
	var twoslices = make([][]int, rows)
	for i := range twoslices {
		twoslices[i] = make([]int, cols)
	}
	fmt.Println(twoslices)
}

func Execute_two_slice() {
	two_slice()
}

package chp1

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func powerSeries2(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

func powerSeries_with_error(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func Exexute_Tuple() {
	var square int
	var cube int
	// returned as tuple
	square, cube = powerSeries(3)
	square2, cube2 := powerSeries2(4)
	square3, cube3, err := powerSeries_with_error(5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Square ", square, "Cube", cube)
	fmt.Println("Square2 ", square2, "Cube2", cube2)
	fmt.Println("Square3 ", square3, "Cube3", cube3)
}

package chp1

import "fmt"

func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}

	if addValue == k {
		num += 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}

	var i int

	for i = 1; i < size; i++ {
		// fmt.Println(" m", m)
		combinations[m] = 1
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l += 1
	}
	return num
}

func Execute_findElementsWithSum() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var addedSum int = 18
	var combinations [19]int
	check := findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
	fmt.Println(check)
	// check2 := findElementsWithSum(arr, 9)
	// fmt.Println(check2)
}

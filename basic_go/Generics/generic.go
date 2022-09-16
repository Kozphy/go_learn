package generic

import "fmt"

func All_generic() {
	fmt.Println("Generics=====================")
	fmt.Println("Type parameters")
	/*
		The type parameters of a function appear between brackets, before the function's arguments.
		func Index[T comparable](s []T, x T) int

		"comparable" is a useful constraint that makes it possible to use the == and != operators on values of the type.
	*/
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	fmt.Println()
	fmt.Println("Generic types")
	var head = List[any]{
		Next: nil,
		Val:  0,
	}

	head.Next = &List[any]{
		Next: nil,
		Val:  "it linklist",
	}

	for n := &head; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
	fmt.Println()
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	Next *List[T]
	Val  T
}

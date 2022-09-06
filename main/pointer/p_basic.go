package pointers

import "fmt"

type Vertex struct {
	X int
	Y int
}

func P_basic() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func Pointers_to_struct() {
	/*
		To access the field X of a struct when we have the struct pointer p
		we could write (*p).X. However, that notation is cumbersome,
		so the language permits us instead to write just p.X,
		without the explicit dereference.
	*/
	v := Vertex{1, 2}
	p := &v
	// can without dereference
	p.X = 1e9
	// optional
	(*p).Y = 8e9
	fmt.Println(v)
}

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}
	v2 = Vertex2{X: 1}
	v3 = Vertex2{}
	p  = &Vertex2{1, 2}
)

func Struct_Literals() {
	fmt.Println(v1, v2, v3, p)
}

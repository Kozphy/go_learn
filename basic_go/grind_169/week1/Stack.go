package week1

type StackVertex struct {
	vertrices []*Vertex
	length    int
}

func (s *StackVertex) IsEmpty() bool {
	return s.length == 0
}

func (s *StackVertex) Push(elem int) {
	vertex := &Vertex{value: elem, visited: true}
	s.length += 1
	s.vertrices = append(s.vertrices, vertex)
}

func (s *StackVertex) Pop() {
	s.vertrices = s.vertrices[:s.length-1]
	s.length -= 1
}

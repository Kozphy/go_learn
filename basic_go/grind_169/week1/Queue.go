package week1

type Vertex struct {
	value   interface{}
	visited bool
}

type QueueVertex struct {
	vertrices []*Vertex
	length    int
}

func (q *QueueVertex) enqueue(element int) {
	NewVertex := &Vertex{value: element, visited: true}
	q.length += 1
	q.vertrices = append(q.vertrices, NewVertex)
}

func (q *QueueVertex) dequeue(element int) {
	q.vertrices = q.vertrices[1:]
	q.length -= 1
}

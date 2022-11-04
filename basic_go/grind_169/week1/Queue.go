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

func (q *QueueVertex) dequeue() {
	q.vertrices = q.vertrices[1:]
	q.length -= 1
}

func (q *QueueVertex) front() Vertex {
	return *q.vertrices[0]
}

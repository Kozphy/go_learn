package json

/*
	source reference: https://go.dev/blog/json
*/
import "encoding/json"

type Message struct {
	Name string
	Body string
	Time int64
}

// Encoding
func Marshal_test(v interface{}) ([]byte, error) {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
}

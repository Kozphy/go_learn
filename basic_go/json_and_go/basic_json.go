package json

/*
	source reference: https://go.dev/blog/json

	Marshal:
	Only data structures that can be represented as valid JSON will be encoded:
	JSON objects only support strings as keys; to encode a Go map type it must be of the form map[string]T (where T is any Go type supported by the json package).

	Channel, complex, and function types cannot be encoded.

	Cyclic data structures are not supported; they will cause Marshal to go into an infinite loop.

	Pointers will be encoded as the values they point to (or ‘null’ if the pointer is nil).

	Unmarshal:
	How does Unmarshal identify the fields in which to store the decoded data?
	For a given JSON key "Foo", Unmarshal will look through the destination struct’s fields to find (in order of preference):
	An exported field with a tag of "Foo".

	An exported field named "Foo", or

	An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".

*/
import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/pkg/errors"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func All_json_and_go() {
	fmt.Println("json and go")
	Marshal_and_UnMarshal_test()

	fmt.Println("UnMarshal_not_find===========")
	UnMarshal_not_find()

	fmt.Println()
	fmt.Println("Generic_json_with_interface========")
	Generic_json_with_interface()

	fmt.Println()
	fmt.Println("Decode_arbitrary_data========")
	Decode_arbitrary_data()

	fmt.Println()
	fmt.Println("Reference_type==================")
	Reference_types()

	fmt.Println()
	fmt.Println("Streaming_encoders_and_decoders==================")
	// json.Streaming_encoders_and_decoders()

}

func Marshal_and_UnMarshal_test() error {

	// Encoding
	m := Message{"Alice", "Hello", 1294706395881547000}
	/*
		If all is well, err will be nil and b will be a []byte containing this JSON data:
		b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	*/
	b, err := json.Marshal(m)
	if err != nil {
		errors.Wrapf(err, "json Marshal fail")
	}

	fmt.Println(b)

	// Decoding
	m = Message{}
	err = json.Unmarshal(b, &m)
	fmt.Println(m)

	return err
}

/*
What happens when the structure of the JSON data doesn’t exactly match the Go type?
Unmarshal will decode only the fields that it can find in the destination type
*/
func UnMarshal_not_find() error {
	b := []byte(`{"Name":"Bob", "Food": "Pickle"}`)
	var m Message
	err := json.Unmarshal(b, &m)
	fmt.Println(m)

	if err != nil {
		errors.Wrapf(err, "json UnMarshal fail")
	}
	return err
}

// Generic JSON with interface
/*
	The json package uses map[string]interface{} and []interface{} values to store arbitrary JSON objects and arrays;
	it will happily unmarshal any valid JSON blob into a plain interface{} value. The default concrete Go types are
	- bool for JSON booleans,

	- float64 for JSON numbers,

	- string for JSON strings, and

	- nil for JSON null.
*/
func Generic_json_with_interface() error {
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	r := i.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)

	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v*2)
	case float64:
		fmt.Println("the reciprocal of i is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}
	return nil
}

// Without knowing this data’s structure, we can decode it into an interface{} value with Unmarshal:
func Decode_arbitrary_data() {
	b := []byte(`{"Name":"Wendnesday", "Age":6, "Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		errors.Wrapf(err, "decode fail")
	}
	fmt.Println(f)

	// To access this data we can use a type assertion to access f’s underlying map[string]interface{}:
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func Reference_types() {
	b := []byte(`{"Name":"Wendnesday", "Age":6, "Parents":["Gomez","Morticia"]}`)
	var m FamilyMember
	/*
		With the var statement we allocated a FamilyMember struct, and then provided a pointer to that value to Unmarshal,
		but at that time the Parents field was a nil slice value. To populate the Parents field,
		Unmarshal allocated a new slice behind the scenes.
		This is typical of how Unmarshal works with the supported reference types (pointers, slices, and maps).
	*/
	err := json.Unmarshal(b, &m)
	if err != nil {
		errors.Wrapf(err, "reference_types err")
	}
	fmt.Println(m)

	/*
		Consider unmarshaling into this data structure:
		type Foo struct {
			Bar *Bar
		}
		If there were a Bar field in the JSON object,
		Unmarshal would allocate a new Bar and populate it.
		If not, Bar would be left as a nil pointer.
	*/
}

/*
Here’s an example program that reads a series of JSON objects from standard input,
removes all but the Name field from each object,
and then writes the objects to standard output:
*/
func Streaming_encoders_and_decoders() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}

		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

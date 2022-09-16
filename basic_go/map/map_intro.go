package mapping

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

func All_map() {
	fmt.Println("Map==============================")
	fmt.Println("Map intro")
	Map_intro()
}

// source article: https://go.dev/blog/maps

// Map signature and introduction
/*
	map[KeyType]ValueType
	KeyType may be any type that is "comparable" (more on this later),
	and ValueType may be any type at all, including another map!
*/

// declare
/*
	declare nil maps behave like an empty map when reading,
	but attempts to write to a nil map will cause a runtime panic.
	var m map[string]int
*/

// init
/*
	Map types are reference types, like pointers or slices,
	and so the value of m above is nil;
	it doesn’t point to an initialized map.


	To initialize a map, use the built in make function:
	var m = make(map[string]int)

	init with empty, identical to using "make" function
	m := map[string]int{}

	The specifics of that data structure are an implementation detail of the runtime
	and are not specified by the language itself.
*/

func Map_intro() {
	// Working with maps
	var m = make(map[string]int)

	// init with empty, identical to using "make" function
	m = map[string]int{}

	m["route"] = 66
	// If the requested key doesn’t exist, we get the value type’s zero value.
	// In this case the value type is int, so the zero value is 0:
	i := m["route"]
	fmt.Println(i)

	// returns on the number of items in a map:
	n := len(m)
	fmt.Println(n)

	// two-value assignment tests for the existence of a key:
	// If that key doesn’t exist, i is the value type’s zero value (0).
	// The second value (ok) is a bool that is "true" if the key exists in the map, and "false" if not.
	_, ok := m["route"]
	if !ok {
		log.Fatal("map error")
	}

	// removes an entry from the map:
	// The delete function doesn’t return anything, and will do nothing if the specified key doesn’t exist.
	delete(m, "route")

	// iterate
	for key, value := range m {
		fmt.Println("Key: ", key, "value: ", value)
	}

	// init with data
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println(commits)

	// Exploiting zero values
	/*
		a map of boolean values can be used as a set-like data structure
		(recall that the zero value for the boolean type is false).
	*/
	// ex1
	// iterate LinkList
	type Node struct {
		Next  *Node
		Value interface{}
	}
	var first *Node

	visited := make(map[*Node]bool)

	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		// The expression visited[n] is true if n has been visited, or false if n is not present.
		visited[n] = true
		fmt.Println(n.Value)
	}

	// ex2
	// filter out elements of Likes of people []*Person and append to "likes" map
	type Person struct {
		Name  string
		Likes []string
	}

	var people []*Person

	people = append(people, []*Person{
		&Person{
			"Braian", []string{"bacon", "cheese"},
		},
		&Person{
			"Kios", []string{"star"},
		},
		&Person{
			"Brita", []string{"cheese"},
		},
	}...)

	likes := make(map[string][]*Person)

	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}

	// print a list of people who like cheese:
	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese.")
	}

	// print the number of people who like bacon:
	fmt.Println(len(likes["bacon"]), "people like bacon")

	// Key types
	/*
		comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.
		Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.

	*/
	// 	Struct can be used to key data by multiple dimensions.
	//  For example, this map of maps could be used to tally web page hits by country
	type Key struct {
		Path, Country string
	}
	hits := make(map[Key]int)
	hits[Key{"/", "vn"}]++

	// Concurrency
	/*
		Maps are not safe for concurrent use: it’s not defined what happens when you read and write to them simultaneously.
		if you need to read from and write to a map from concurrently executing goroutines,
		One common way to protect maps is with "sync.RWMutex".
	*/
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	// To read from counter, take the read lock:
	counter.RLock()
	nw := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key: ", nw)

	// To write to the counter, take the write lock:
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	counter.RLock()
	nw = counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("after write to counter: ", nw)

	// Iteration order
	/*
		When iterating over a map with a range loop, the iteration order is not specified and
		is not guaranteed to be the same from one iteration to the next.
		If you require a stable iteration order you must maintain a separate data structure that specifies that order.
	*/
	m1 := map[int]string{
		1:   "some_key",
		2:   "lo2",
		3:   "h2o",
		5:   "swea",
		4:   "woi",
		100: "twik",
	}
	var keys []int
	for k := range m1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key: ", k, "Value:", m1[k])
	}

}

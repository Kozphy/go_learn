package exercise

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func exercise_stringers() {
	fmt.Println("Exercise: Stringers")

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println()
	fmt.Println("Errors")
	/*
		"error" type is a built-in interface similar to fmt.Stringer
		type error interface {
			Error() string
		}
	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var res = make([]string, len(ip))
	for i, add := range ip {
		res[i] = strconv.Itoa(int(add))
	}
	return strings.Join(res, ".")
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

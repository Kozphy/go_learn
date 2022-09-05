package greetings

import (
	"regexp"
	"testing"
)

/*
Test function names have the form TestName, where Name says something about the specific test.
test functions take a pointer to the testing package's testing.T type as a parameter.

run "go test" to executes the test in the directory where you want to test
*/

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "",  error`, msg, err)
	}
}

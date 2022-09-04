package main

import (
	"fmt"
	"log"

	"example.com/hello"
	functions "github.com/zixas/go_learn/func"
	"github.com/zixas/go_learn/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")

	fmt.Println("Hello world")
	fmt.Println(quote.Go())

	var message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	hello.Hello()
	functions.Test()
}

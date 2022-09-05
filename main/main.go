package main

import (
	"fmt"
	"log"

	"example.com/hello"
	functions "github.com/zixas/go_learn/func"
	"github.com/zixas/go_learn/greetings"
	mapping "github.com/zixas/go_learn/map"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")

	fmt.Println("Hello world")
	fmt.Println(quote.Go())

	hello.Hello()
	functions.Test()

	names := []string{"Gladys", "Samantha", "Darrin"}
	var message, err = greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	mapping.Map_intro()

}

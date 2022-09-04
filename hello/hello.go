package hello

import (
	"fmt"
	"github.com/zixas/go_learn/greetings"
)

func Hello() {
	message, _ := greetings.Hello("Gladys")
	fmt.Println(message)
}
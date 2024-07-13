package main

import (
	"fmt"
	"github.com/kolewttd/go/greetings"
)

func main() {
	message := greetings.Hello("wtt")
	fmt.Println(message)
}

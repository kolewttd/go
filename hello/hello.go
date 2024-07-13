package main

import (
	"fmt"
	"github.com/kolewttd/greetings"
)

func main() {
	message := greetings.Hello("wtt")
	fmt.Println(message)
}

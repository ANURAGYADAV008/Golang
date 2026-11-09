package main

import "fmt"

func deferkaro() {
	message := "hello"
	greetings := func(name string) {
		fmt.Println(message + name)
	}

	defer greetings("Alice")
	defer greetings("BOB")
	//message = "HI"

}

package main

import "fmt"

func clousers() {
	okk := counter()
	fmt.Println(okk())
	fmt.Println(okk())

}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Hello pointers, Motherfucker!")

	foo := 23
	println(foo)

	pointerFoo := &foo
	println(pointerFoo)

	println(*pointerFoo)
	*pointerFoo = 10

	println(foo)
}

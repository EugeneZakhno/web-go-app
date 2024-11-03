package main

import "fmt"

func main() {
	fmt.Println("Hello, Motherfuckers!")

	foo := 23
	println(foo)

	pointerFoo := &foo
	println(pointerFoo)

	println(*pointerFoo)
	*pointerFoo = 10

	println(foo)

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

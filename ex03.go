package main

import "fmt"

func main() {
	var a, b int = 1, 2
	fmt.Println("a is", a, "and b is", b)

	a = 10
	fmt.Println("Now a is", a)

	var c = 10 // inferred type
	fmt.Println("c is", c)

	d := "Hello world" // like this you don't need type neither var
	fmt.Println("d says", d)
}

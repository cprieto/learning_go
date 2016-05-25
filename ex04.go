package main

import "fmt"

func main() {
	fmt.Println("** Using for **")

	// It will do-it until condition is true
	i := 0
	for i <= 5 {
		i++
	}
	fmt.Println("Finished the first loop")

	for j := 0; j <= 5; j++ {
		// This is the classic loop
		fmt.Print(".")
	}
	fmt.Println("Finished second loop")

	for {
		fmt.Println("Eternal loop until I break")
		break
	}
}

package main

import "fmt"

// * dereference the value at pointer location and update it with new value
func updateName(n *string) {
	*n = "updated value"
}

func main() {

	x := "hey there"

	// updateName(x)
	// fmt.Println("Location of x in memory", &x)

	m := &x

	// fmt.Println("Value at memory address", *m)

	updateName(m)

	fmt.Println(x)

}

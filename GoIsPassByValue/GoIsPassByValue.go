package main

import "fmt"

func updateVal(n string) string {
	n = "de"
	return n
}

func updateMap(x map[string]float64) {
	x["coffee"] = 2.99
}

func main() {

	k := "hello"

	k = updateVal(k)
	fmt.Println(k)

	// strings, ints, floats, booleans, arrays and structs are stored without pointer

	// maps, slices, and functions are stored via a pointer wrapper to the location in memory
	menu := map[string]float64{
		"hello":  32.344,
		"sdffef": 44.43,
	}

	updateMap(menu)
	fmt.Println(menu)

}

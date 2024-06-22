package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":         4.99,
		"pie":          7.99,
		"sald":         6.99,
		"toffe powder": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps

	for k, v := range menu {
		fmt.Println(k, v)
	}

	// ints as key type
	phonebook := map[int]string{
		26783232: "mario",
		54544545: "hitesh",
		53476654: "deepak",
	}

	fmt.Println(phonebook[26783232])
	phonebook[54544545] = "jumar"
	fmt.Println(phonebook)

	phonebook[26783232] = "kk"
	fmt.Println(phonebook)

}

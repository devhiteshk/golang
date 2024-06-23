package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	// strings are stored as utf-8 encodings so the length wont work
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

}

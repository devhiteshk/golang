package main

import (
	"fmt"
	"math"
	"strings"
)

func greeting(n string) {
	fmt.Println("Hello", n)
}

func sayBye(n string) {
	fmt.Println("GoodBye", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {

	// greeting("Hitesh")
	// sayBye("Deepak")
	// cycleNames([]string{"hitesh", "deepak"}, greeting)
	// cycleNames([]string{"hitesh", "deepak"}, sayBye)

	// a1 := circleArea(3.14137)
	// a2 := circleArea(10.14137)
	// fmt.Println(a1)
	// fmt.Println(a2)

	fn, sn := getInitials("Hitesh Kumar")
	fmt.Println(fn, sn)

	fn1, sn1 := getInitials("Hitesh")
	fmt.Println(fn1, sn1)
}

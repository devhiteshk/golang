package main // tells go to create a standalone executable

import (
	"fmt"
	"sort"
	"strings"
)

// var someName := "hello"  // --> cant decalre like this outside of function block

func main() {
	fmt.Println("hello world")
	var nameOne string = "hello"

	var nameTwo = "hello2"

	nameFour := "hello"

	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// intrgers
	var ageOne int = 20
	var ageTwo = 22
	ageThree := 34

	fmt.Println(ageTwo, ageOne, ageThree)

	// bits and memory
	// var numOne int8 = 25
	// var numTwo int16 = 129
	// var numThree uint8 = 25

	age := 23
	name := "hitesh"

	fmt.Print("Hello, \n")
	fmt.Print("world, \n")
	fmt.Print("new line \n")

	fmt.Println("Hey")
	fmt.Println("mt age is", age, "and my name is", name)

	// formatted string %_ = format specifier
	// %v -> variables
	// %q -> string %t
	// %T -> Type of variable
	// %f -> floating point variables
	// %x.xf -> floating point variables with number of digits after decimal
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("my age is %T and my name is %T \n", age, name)
	fmt.Printf("you scored %f \n", 255.55)
	fmt.Printf("you scored %0.1f \n", 255.55)

	// fixed length array
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"youshi", "mario", "peach", "bowser"}

	names[1] = "luigi"

	fmt.Println("array", ages)
	fmt.Println("array", names)

	// slice (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 23
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:4]
	rangeTwo := names[2:]
	rangethree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangethree)

	// strings std lib

	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))

	fmt.Println("original string value", greeting)

	fmt.Println(strings.Index(greeting, "th"))

	fmt.Println(strings.Split(greeting, " "))

	ages2 := []int{45, 35, 33, 20, 23, 1, 2, 6, 10, 11}

	sort.Ints(ages2)
	fmt.Println(ages2)

	index := sort.SearchInts(ages2, 30)
	fmt.Println(index)

	namesx := []string{"hi", "me", "a", "ab", "x"}
	sort.Strings(namesx)
	fmt.Println(namesx)

	fmt.Println(sort.SearchStrings(namesx, "x"))
}

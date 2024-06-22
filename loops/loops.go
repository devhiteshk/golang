package main

import "fmt"

func main() {

	for x := 0; x < 10; x++ {
		fmt.Println("value of x is:", x)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"mario", "luigui", "youshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new"
	}

	fmt.Println(names)

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age == 45)
	fmt.Println(age > 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // break
		}
		fmt.Println("index, value", index, value)
	}

}

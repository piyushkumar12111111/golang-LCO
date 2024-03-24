package main

import "fmt"

func main() {

	var number int = 10

	var  result string

	if number > 10 {
		result = "Greater"
	} else if number == 10 {
		result = "Equal"
	} else {
		result = "Lesser"
	}

	fmt.Println("Result is: ", result)

//! if else with init statement 
	if number%2 ==0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if checking := 3 ; checking < 10 {
		fmt.Println("This number is less than 10")
	} else {
		fmt.Println("This number is greater than 10")
	}
}

package main

import "fmt"

func main() {
	fmt.Println("go lang for loop")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	for i := 0; i < len(days); i++ {
		fmt.Println("Day is: ", days[i])
	}

	//! for loop with range
	for index, day := range days {
		fmt.Printf("Day %v is: %v\n", index, day)
	}

	//! if you don't want to use index then use _ instead of index

	//! most used
	for _, day := range days {
		fmt.Println("Day is: ", day)
	}
	//! for while loop
	roughvalue := 1

	for roughvalue < 10 {

		if roughvalue == 5 {

			break
		}

		fmt.Println("Rough value is: ", roughvalue)
		roughvalue++
	}

}

//! there is only for loop in go lang no while loop and do while loop

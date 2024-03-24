package main

import "fmt"

func main() {

	fmt.Println("go lang array check")

	var fruitlist [2]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))


	var veglist = [3]string{"potato", "tomato", "onion"}

	fmt.Println("Veg list is: ", veglist)
	fmt.Println("Veg list is: ", len(veglist))
}

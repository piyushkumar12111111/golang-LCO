
package main

func main(){


	number := 10

	var ptr = &number

	println("Value of number is: ", ptr)
	println("Address of number is: ", *ptr)

	*ptr = *ptr +20

	println("Value of number is: ", *ptr)
}
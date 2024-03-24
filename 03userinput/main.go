package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is variable example\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")

	//! _ is used to ignore the error value also called as COMMA OK IDIOM
	input, _ := reader.ReadString('\n')

	fmt.Println("Text entered: ", input)

}

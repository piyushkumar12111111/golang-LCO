package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome pizza app\n")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Your order is: ", input)

	numorder, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {

		fmt.Println(err)
	} else {

		fmt.Println("Your order is: ", numorder+1)
	}

}

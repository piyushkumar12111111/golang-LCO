package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Hello, World!\n")

	content := " This is my file piyushio.com"

	file, err := os.Create("myfile.txt")

	if err != nil {
		panic(err)
		return
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
		file.Close()
		return
	}
	fmt.Println("Length of the file is: ", length)
	defer file.Close() //! defer is used to close the file after the main function is executed

	//! calling function to read the file
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Text file data is: ", string(databyte))
}

// ! checknilerror is used to check the error
func checknilerror(err error) {
	if err != nil {
		panic(err)
	}
}

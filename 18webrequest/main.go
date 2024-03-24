package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Get Request Web Server")

	PerformGetRequest()

}

func PerformGetRequest() {

	//	fmt.Println("Perform Get Request")

	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {

		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(" Status Code: ", response.StatusCode)
	fmt.Println(" contentLength: ", response.ContentLength)

	var responseString strings.Builder

	content, err := ioutil.ReadAll(response.Body) //! read the content of the response it is in bytes format
	bytecount, err := responseString.Write(content)

	fmt.Println("Byte Count: ", bytecount)
	fmt.Println(responseString.String()) //! second method of handling data using String method better method

	if err != nil {
		panic(err)
	}

	data := string(content)

	fmt.Println("Data is: ", data)
}

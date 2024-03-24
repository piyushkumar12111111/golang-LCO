package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {

	fmt.Println("URL is: ", myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Result is: ", result)
	}

	//! url parssing information
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//! Query parsing
	queryParams := result.Query()
	fmt.Printf("Type of the Query Params is: %T\n", queryParams)

	fmt.Println(queryParams["coursename"])

	for _, val := range queryParams {
		fmt.Println("Param is: ", val)
	}

	//! URL building formation of url from parts

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}

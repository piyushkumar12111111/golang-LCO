// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// const url = "https://lco.dev"

// func main() {

// 	fmt.Println("go web request check")

// 	response, err := http.Get(url)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Response is of Type: %T", response)

// 	//fmt.Println("Response value is: %v", response)

// 	databytes ,err = ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	  fmt.Println("Data is: ", string(databytes))

// 	defer response.Body.Close()

// }

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}

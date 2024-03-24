package main

import (
	"encoding/json"
	"fmt"
)

type cources struct {
	Name     string `json:"cource_name"`
	Price    int
	Platform string
	Password string `json:"-"` //! this will not be shown in json response
	Tags     []string
}

func main() {

	fmt.Println("Json response handling")

	//encodeJson()

	DecodeJson()

}

func encodeJson() {

	lcocources := []cources{

		{Name: "React Native", Price: 299, Platform: "Udemy", Password: "password", Tags: []string{"web-dev", "mobile-dev"}},
		{Name: "Angular", Price: 199, Platform: "Udemy", Password: "password", Tags: []string{"web-dev"}},
		{Name: "MERN", Price: 199, Platform: "Udemy", Password: "password", Tags: []string{"web-dev"}},
		{Name: "LCO", Price: 0, Platform: "LearnCodeOnline", Password: "password", Tags: []string{"web-dev", "mobile-dev"}},
	}

	//! json.Marshal() will convert the struct into json
	//! json.MarshalIndent() will convert the struct into json with proper indentation
	//! json.Unmarshal() will convert the json into struct
	//! json.NewEncoder() will convert the struct into json and write it to the writer
	//! json.NewDecoder() will convert the json into struct and read it from the reader

	finalJsonmarsh, err := json.Marshal(lcocources)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJsonmarsh))

	//! for proper output json use MarshalIndent

	finalJson, err := json.MarshalIndent(lcocources, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

	//fmt.Println("%s\n", finalJson)
}

func DecodeJson() {

	jsonDataFromWeb := []byte(

		`{
			"cource_name": "React Native",
			"Price": 299,
			"Platform": "Udemy",
			"Tags": [
					"web-dev",
					"mobile-dev"
			]
	}`,
	)

	var mycource cources

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &mycource)
		fmt.Println(mycource)
	} else {
		fmt.Println("Json is not valid")
	}

	//! convert to key value pairing

	var myonlinedata map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myonlinedata)
	fmt.Println(myonlinedata)


	for key, value := range myonlinedata {
		fmt.Printf("Key is: %v and value is: %v\n", key, value)
	}
}

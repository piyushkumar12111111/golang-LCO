package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println(languages)

	delete(languages, "RB")

	fmt.Println(languages["JS"])

 //! for loop to iterate over map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

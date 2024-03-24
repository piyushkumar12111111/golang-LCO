package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitlist = []string{"Apple", "Mango", "Banana"}

	//fmt.Println(" Type of Fruit list is  %T\n", fruitlist)

	fruitlist = append(fruitlist, "Grapes")

	//fruitlist.append(fruitlist[1:2])

	fmt.Println(fruitlist)

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 465
	highscore[3] = 345

	highscore = append(highscore, 555, 666, 777, 888, 999, 000)
	sort.Ints(highscore)

	fmt.Println(highscore)

	var cources = []string{"react", "angular", "vue", "svelte"}

	fmt.Println(cources)

	var index int = 2

	//! remove the element from the slice using append last in [index+1:] not included 

	cources = append(cources[:index], cources[index+1:]...)

	fmt.Println(cources)


	

}

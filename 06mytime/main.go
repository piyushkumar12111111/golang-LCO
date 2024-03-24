package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcom to time zone\n")

	presentTime := time.Now()

	fmt.Println(presentTime)

	createTime := time.Date(2020, time.January, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createTime)

	fmt.Println(createTime.Format("01-02-2006 15:04:05 Monday"))

}

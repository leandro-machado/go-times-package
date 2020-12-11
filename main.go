package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()

	fmt.Println(date.Format("02/Jan/2006 15:04:05"))

	fmt.Println(date.Format("January/2006"))

	fmt.Println(date.Format("Jan/2006"))

	fmt.Println(date.Format("Monday 02/January/2006 15:04:05"))

	location, _ := time.LoadLocation("America/Vancouver")
	fmt.Println(date.In(location))
}

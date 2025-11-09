package main

import (
	"fmt"
	"time"
)

func greetingTimeDate() {
	greeting := "Hello, "
	currentTime := time.Now()
	dateFormat := currentTime.Format("01/02/2006")
	timeFormat := currentTime.Format("03:04 PM")
	fmt.Println(greeting + "the date is " + dateFormat + ", and the time is " + timeFormat + ".")
}

func main() {
	greetingTimeDate()
}

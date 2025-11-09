package main

import (
        "fmt"
        "time"
)

func greetingTimeDate() {
        greeting := "Hello!"
        currentTime := time.Now()
        dateFormat := currentTime.Format("01/02/2006")
        timeFormat := currentTime.Format("03:04PM")
        fmt.Println(greeting + ", the current date is " + dateFormat + ", and the current UTC time is " + timeFormat + ".")
}

func main() {
        greetingTimeDate()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 100 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Println("Duration : %d\n", duration3)
}

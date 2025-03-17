package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Mohammad Nawawi", "Nawawi"))
	fmt.Println(strings.Split("Mohammad Nawawi", " "))
	fmt.Println(strings.ToLower("Mohammad Nawawi"))
	fmt.Println(strings.ToUpper("Mohammad Nawawi"))
	fmt.Println(strings.Trim(" Mohammad Nawawi ", " "))
	fmt.Println(strings.ReplaceAll("Mohammad Nawawi", "Nawawi", "Naw"))
}

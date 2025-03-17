package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "Database username")
	var password *string = flag.String("password", "root", "Database password")
	var hostname *string = flag.String("hostname", "localhost", "Database hostname")
	var port *int = flag.Int("port", 3306, "Database port")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Hostname:", *hostname)
	fmt.Println("Port:", *port)
}
